package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/kechako/mktree/node"
)

func _main() (int, error) {
	args := os.Args[1:]

	var r io.Reader
	if len(args) > 0 {
		file, err := os.Open(args[0])
		if err != nil {
			return 1, fmt.Errorf("could not open the file : %s", args[0])
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
		return 2, fmt.Errorf("scan error : %v", s.Err())
	}

	if root == nil {
		return 0, nil
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	root.Print(w)

	return 0, nil
}

func main() {
	code, err := _main()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(code)
	}
}
