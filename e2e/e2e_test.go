package e2e_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/while1malloc0/docdir/cmdtest"
)

func TestMain(m *testing.M) {
	cleanup, err := setup()
	if err != nil {
		fmt.Printf("unexpected error in e2e setup: %v\n", err)
		os.Exit(1)
	}
	exitCode := m.Run()
	if err := cleanup(); err != nil {
		fmt.Printf("unexpected error in e2e cleanup: %v\n", err)
		os.Exit(1)
	}
	os.Exit(exitCode)
}

func TestE2E(t *testing.T) {
	cases := []struct {
		description string
		testPath    string
	}{
		{"a directory with no subdirectories prints the directory", "e2e/testdata/simple.ct"},
		{"a directory with subdirectories prints a-la tree", "e2e/testdata/nested.ct"},
		{"a directory with description files prints the descriptions", "e2e/testdata/nested-description.ct"},
		{"mixing described and nondescript directories prints both", "e2e/testdata/mixed-description.ct"},
		{"the -skip-missing flag skips directories without descriptions", "e2e/testdata/skipped-missing.ct"},
	}
	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			run, cleanup, err := cmdtest.ReadFile(tc.testPath)
			if err != nil {
				t.Fatalf("unexpected error reading test data: %v", err)
			}
			run(t)
			if err := cleanup(); err != nil {
				t.Fatalf("unexpected error cleaning up after execution: %v", err)
			}
		})
	}
}

func setup() (func() error, error) {
	// cd into project root so all e2e tests assume same dir
	if err := os.Chdir(".."); err != nil {
		return nil, err
	}

	// build test version of binary
	cmd := exec.Command("go", "build", "-o", "dist/test/docdir", "./cmd/main.go")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%s", stderr.String())
	}

	// add test version of binary to PATH
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dist := filepath.Join(cwd, "dist/test")
	os.Setenv("PATH", fmt.Sprintf("%s:%s", os.Getenv("PATH"), dist))

	return func() error {
		// remove test version of binary
		// NB: no need to reset PATH, go tests are run in a subprocess, so PATH
		// changes are temporary
		return os.RemoveAll(dist)
	}, nil
}
