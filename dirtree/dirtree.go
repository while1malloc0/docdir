// package dirtree implements a tree structure of directories
package dirtree

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const DefaultDescriptionFile = "DESCRIPTION"

type Node struct {
	Name        string
	Description string
	Children    []*Node
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
	descriptionData, err := os.ReadFile(filepath.Join(path, DefaultDescriptionFile))
	if err == nil {
		n.Description = string(descriptionData)
	}

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
	return align(s.String())
}

// based on https://github.com/campoy/tools/tree
func visit(w io.Writer, node *Node, prefix string) error {
	s := node.Name
	if node.Description != "" {
		s = fmt.Sprintf("%s # %s", node.Name, node.Description)
	}
	fmt.Fprintln(w, s)
	padding := "│   "
	for i, child := range node.Children {
		if i == len(node.Children)-1 {
			fmt.Fprintf(w, prefix+"└── ")
			padding = "    "
		} else {
			fmt.Fprintf(w, prefix+"├── ")
		}
		err := visit(w, child, prefix+padding)
		if err != nil {
			return err
		}
	}
	return nil
}

func align(in string) string {
	var out strings.Builder

	longest := findLongest(in)
	// longest name+indent pair + 2 spaces
	lengthToMeet := longest + 2
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "#")
		if len(parts) == 1 {
			// no description provided
			fmt.Fprintln(&out, line)
			continue
		}
		nameAndIndent := strings.TrimRight(parts[0], " ")
		description := strings.TrimSpace(parts[1])
		paddingNeeded := lengthToMeet - normalizedLen(nameAndIndent)
		padding := strings.Repeat(" ", paddingNeeded)
		fmt.Fprintf(&out, "%s%s# %s\n", nameAndIndent, padding, description)
	}
	return out.String()
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

// return the length of a string from a visual perspective, e.g. each character
// is one space, regardless of unicode
func normalizedLen(s string) int {
	s = strings.Replace(s, "├──", "   ", 1)
	s = strings.Replace(s, "└──", "   ", 1)
	return len(s)
}

// returns the length of the longest string after stripping descriptions and
// normalizing string len
func findLongest(s string) int {
	scanner := bufio.NewScanner(strings.NewReader(s))
	longest := -1
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "#")
		nameAndIndent := strings.TrimRight(parts[0], " ")
		visualLen := normalizedLen(nameAndIndent)
		if visualLen > longest {
			longest = visualLen
		}
	}
	return longest
}
