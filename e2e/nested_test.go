package e2e_test

import (
	"testing"

	"github.com/while1malloc0/docdir/cmdtest"
)

func TestNested(t *testing.T) {
	run, cleanup, err := cmdtest.ReadFile("e2e/testdata/nested.ct")
	if err != nil {
		t.Fatalf("unexpected error reading test data: %v", err)
	}
	run(t)
	if err := cleanup(); err != nil {
		t.Fatalf("unexpected error cleaning up after execution: %v", err)
	}
}
