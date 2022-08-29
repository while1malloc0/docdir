package e2e_test

import (
	"testing"

	"github.com/while1malloc0/docdir/cmdtest"
)

func TestDescription(t *testing.T) {
	run, cleanup, err := cmdtest.ReadFile("e2e/testdata/nested-description.ct")
	if err != nil {
		t.Fatalf("unexpected error reading test data: %v", err)
	}
	run(t)
	if err := cleanup(); err != nil {
		t.Fatalf("unexpected error cleaning up test: %v", err)
	}
}
