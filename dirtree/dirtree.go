// package dirtree implements a tree structure of directories
package dirtree

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Node struct {
	Name     string
	Children []*Node
}

func New(path string) (*Node, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, nil
	}

	n := Node{Name: fi.Name(), Children: []*Node{}}

	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	_ = dir.Close()

	names = filterNonDirs(path, names)
	sort.Strings(names)
	for _, name := range names {
		child, err := New(filepath.Join(path, name))
		if err != nil {
			return nil, err
		}
		n.Children = append(n.Children, child)
	}

	return &n, nil
}

func (n *Node) String() string {
	var s strings.Builder
	visit(&s, n, "")
	return s.String()
}

// based on https://github.com/campoy/tools/tree
func visit(w io.Writer, node *Node, indent string) error {
	fmt.Fprintln(w, node.Name)
	add := "│   "
	for i, child := range node.Children {
		if i == len(node.Children)-1 {
			fmt.Fprintf(w, indent+"└── ")
			add = "    "
		} else {
			fmt.Fprintf(w, indent+"├── ")
		}
		err := visit(w, child, indent+add)
		if err != nil {
			return err
		}
	}
	return nil
}

func filterNonDirs(path string, candidates []string) []string {
	var dirs []string
	for _, candidate := range candidates {
		fi, _ := os.Stat(filepath.Join(path, candidate))
		if fi.IsDir() {
			dirs = append(dirs, candidate)
		}
	}
	return dirs
}