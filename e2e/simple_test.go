package e2e_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	testData, err := os.Open("e2e/testdata/simple.ct")
	if err != nil {
		t.Fatalf("could not read test data: %v", err)
	}
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
