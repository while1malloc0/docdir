// package dirtree implements a tree structure of directories
package dirtree

import (
	"os"
	"path/filepath"
	"sort"
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
