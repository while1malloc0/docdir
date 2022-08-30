package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/while1malloc0/docdir/dirtree"
)

var (
	skipMissing = flag.Bool("skip-missing", false, "skip directories missing description files (and their subdirectories)")
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
	root, err := dirtree.New(path, *skipMissing)
	if err != nil {
		return err
	}
	fmt.Print(root.String())
	return nil
}
