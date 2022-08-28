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

func TestChildren(t *testing.T) {
	root, err := dirtree.New("testdata/a")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	children := root.Children
	if len(children) != 2 {
		t.Fatalf("expected node A to have 2 children, had %d", len(children))
	}

	if children[0].Name != "aa" {
		t.Fatalf("expected first child of node A to have name aa, had %s", children[0].Name)
	}

	if children[1].Name != "bb" {
		t.Fatalf("expected second child of node A to have name bb, had %s", children[1].Name)
	}
}

func TestGrandChildren(t *testing.T) {
	root, err := dirtree.New("testdata/a")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	grandchildren := root.Children[0].Children
	if len(grandchildren) != 3 {
		t.Fatalf("expected node A to have 3 grandchildren from node AA, had %d", len(grandchildren))
	}

	if grandchildren[0].Name != "aaa" {
		t.Fatalf("expected first child of node AA to have name aaa, had %s", grandchildren[0].Name)
	}

	if grandchildren[1].Name != "bbb" {
		t.Fatalf("expected second child of node AA to have name bbb, had %s", grandchildren[1].Name)
	}

	if grandchildren[2].Name != "ccc" {
		t.Fatalf("expected third child of node AA to have name ccc, had %s", grandchildren[2].Name)
	}
}
