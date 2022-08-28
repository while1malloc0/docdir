package dirtree_test

import (
	"testing"

	"github.com/while1malloc0/docdir/dirtree"
)

func TestRead(t *testing.T) {
	got, err := dirtree.New("testdata/a")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.Name != "a" {
		t.Fatalf("expected name 'a', got %v", got.Name)
	}
}
