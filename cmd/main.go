package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
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
	return visit(path, "")
}

// based on https://github.com/campoy/tools/tree
func visit(path, indent string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("stat %s: %v", path, err)
	}
	if !fi.IsDir() {
		// skip all non-directories
		return nil
	}
	fmt.Println(fi.Name())

	dir, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open %s: %v", path, err)
	}
	names, err := dir.Readdirnames(-1)

	names = filterNonDirs(path, names)

	_ = dir.Close() // safe to ignore this error.
	if err != nil {
		return fmt.Errorf("read dir names %s: %v", path, err)
	}
	// names = removeHidden(names)

	sort.Strings(names)
	add := "│   "
	for i, name := range names {
		if i == len(names)-1 {
			fmt.Printf(indent + "└── ")
			add = "    "
		} else {
			fmt.Printf(indent + "├── ")
		}
		err := visit(filepath.Join(path, name), indent+add)
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
