package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kechako/mktree/node"
)

var (
	lineVerticalAndRight = "\u251c"
	lineUpAndRight       = "\u2514"
	lineHorizontal       = "\u2500"
	lineVertical         = "\u2502"
	lineSpace            = "\u00a0"
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

	printNode(root)

	return 0, nil
}

func printNode(n *node.Node) {
	var texts []string
	texts = append(texts, n.Text())

	if n.Parent() == nil {
		// Do nothing
	} else if n.Next() != nil {
		texts = append(texts, " ")
		texts = append(texts, lineHorizontal)
		texts = append(texts, lineHorizontal)
		texts = append(texts, lineVerticalAndRight)
	} else {
		texts = append(texts, " ")
		texts = append(texts, lineHorizontal)
		texts = append(texts, lineHorizontal)
		texts = append(texts, lineUpAndRight)
	}

	for p := n.Parent(); p != nil; p = p.Parent() {
		if p.Parent() == nil {
			// Do nothing
		} else if p.Next() != nil {
			texts = append(texts, " ")
			texts = append(texts, lineSpace)
			texts = append(texts, lineSpace)
			texts = append(texts, lineVertical)
		} else {
			texts = append(texts, " ")
			texts = append(texts, " ")
			texts = append(texts, " ")
			texts = append(texts, " ")
		}
	}

	reverse(texts)
	fmt.Println(strings.Join(texts, ""))

	for c := n.FirstChild(); c != nil; c = c.Next() {
		printNode(c)
	}
}

func reverse(t []string) {
	l := len(t)
	mid := l / 2
	for i := 0; i < mid; i++ {
		t[i], t[l-i-1] = t[l-i-1], t[i]
	}
}

func main() {
	code, err := _main()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(code)
	}
}
