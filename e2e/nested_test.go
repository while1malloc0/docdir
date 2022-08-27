package e2e_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestNested(t *testing.T) {
	originalDir, _ := os.Getwd()
	testData, err := os.Open("e2e/testdata/nested.ct")
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
	os.Chdir(originalDir)
}
