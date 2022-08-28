package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/while1malloc0/docdir/dirtree"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()
	path := flag.Arg(0)
	root, err := dirtree.New(path)
	if err != nil {
		return err
	}
	return visit(root, "")
}

// based on https://github.com/campoy/tools/tree
func visit(node *dirtree.Node, indent string) error {
	fmt.Println(node.Name)
	add := "│   "
	for i, child := range node.Children {
		if i == len(node.Children)-1 {
			fmt.Printf(indent + "└── ")
			add = "    "
		} else {
			fmt.Printf(indent + "├── ")
		}
		err := visit(child, indent+add)
		if err != nil {
			return err
		}
	}
	return nil
}