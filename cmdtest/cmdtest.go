// cmdtest is a library for running tests against CLIs.
// It's based on Google's go-cmdtest library, and uses the same syntax.
// However, there are a few notable differences:
//
// 1. go-cmdtest creates a temporary directory to execute in. This package
// requires the caller to manage its calling directory. The reasoning behind
// that decision is to allow tests to do things like assert against their
// examples folder without needing to copy those directories in the test.
//
// 2. go-cmdtest allows for running Go functions as if they were compiled
// binaries. This package does not, mostly because there's not really a need for
// it in testing this program.
//
// 3. go-cmdtest treats any non-zero exit code as a failure unless explicitly
// instructed otherwise. This package just compares the content of std{out,err}
//
// 4. go-cmdtest allows for assertions anywhere in the .ct file. This package
// assumes that each .ct file will have some number of commands run, and then a
// single output to be asserted against.
//
// Because all execution happens in the same process, cmdtest is not safe to be
// run in parallel.
//
// See https://github.com/google/go-cmdtest for information on the .ct syntax
package cmdtest

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

type ExecFunc func(t *testing.T)
type CleanupFunc func() error

func ReadFile(path string) (ExecFunc, CleanupFunc, error) {
	originalDir, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	testData, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	execFunc := func(t *testing.T) {
		s := bufio.NewScanner(testData)
		var got strings.Builder
		var want strings.Builder
		for s.Scan() {
			line := s.Text()
			switch {
			case strings.HasPrefix(line, "#"): // comment, skip
				continue
			case strings.HasPrefix(line, "$"): // command, execute
				execCommand(t, line, &got)
			default:
				fmt.Fprintln(&want, line)
			}
		}
		if got.String() != want.String() {
			t.Fatalf("expected:\n%s\ngot:\n%s\n", want.String(), got.String())
		}
	}

	cleanupFunc := func() error {
		return os.Chdir(originalDir)
	}

	return execFunc, cleanupFunc, nil
}

func execCommand(t *testing.T, line string, out io.Writer) {
	t.Helper()
	line = line[1:] // trim off the $
	var cmdAndArgs []string
	for _, part := range strings.Split(line, " ") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		cmdAndArgs = append(cmdAndArgs, part)
	}
	switch cmdAndArgs[0] { // fake shell built-ins
	case "cd":
		err := os.Chdir(cmdAndArgs[1])
		if err != nil {
			t.Fatalf("unexpected error running command cd: %v", err)
		}
	default:
		cmd := exec.Command(cmdAndArgs[0], cmdAndArgs[1:]...)
		cmd.Stdout = out
		cmd.Stderr = out
		cmd.Env = nil

		if err := cmd.Run(); err != nil {
			t.Fatalf("unexpected error running command: %v", err)
		}
	}
}
