// package dirtree implements a tree structure of directories
package dirtree

import "os"

type Node struct {
	Name     string
	Children []Node
}

func New(path string) (*Node, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, nil
	}
	return &Node{Name: fi.Name()}, nil
}
