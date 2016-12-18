package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"

	"github.com/kechako/mktree/node"
)

var (
	lineBranching = "\u251c\u2500\u2500\u0020"
	lineTerminal  = "\u2514\u2500\u2500\u0020"
	lineVertical  = "\u2502\u00a0\u00a0\u0020"
	lineSpaces    = "\u0020\u0020\u0020\u0020"
)

func _main() (int, error) {
	args := os.Args[1:]

	var r io.Reader
	if len(args) > 0 {
		file, err := os.Open(args[0])
		if err != nil {
			return 1, fmt.Errorf("Could not open the file : %s", args[0])
		}
		defer file.Close()
		r = file
	} else {
		r = os.Stdin
	}

	var root *node.Node
	var current *node.Node
	s := bufio.NewScanner(r)
	for s.Scan() {
		n := node.New(s.Text())
		if n.IsEmpty() {
			continue
		}

		if root == nil {
			root = n
		} else {
			p := current.Add(n)
			if p == nil {
				break
			}
		}

		current = n
	}
	if s.Err() != nil {
		return 2, fmt.Errorf("Scan error : %v", s.Err())
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	printNode(root, w)

	return 0, nil
}

func printNode(n *node.Node, w io.Writer) {
	outputs := list.New()

	outputs.PushBack(n.Text() + "\n")

	if n.Parent() == nil {
		// Do nothing
	} else if n.Next() != nil {
		outputs.PushBack(lineBranching)
	} else {
		outputs.PushBack(lineTerminal)
	}

	for p := n.Parent(); p != nil; p = p.Parent() {
		if p.Parent() == nil {
			// Do nothing
		} else if p.Next() != nil {
			outputs.PushBack(lineVertical)
		} else {
			outputs.PushBack(lineSpaces)
		}
	}

	for e := outputs.Back(); e != nil; e = e.Prev() {
		fmt.Fprint(w, e.Value)
	}

	for c := n.FirstChild(); c != nil; c = c.Next() {
		printNode(c, w)
	}
}

func main() {
	code, err := _main()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(code)
	}
}
