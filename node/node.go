package node

import (
	"container/list"
	"fmt"
	"io"
	"strings"
)

var (
	lineBranching = "\u251c\u2500\u2500\u0020"
	lineTerminal  = "\u2514\u2500\u2500\u0020"
	lineVertical  = "\u2502\u00a0\u00a0\u0020"
	lineSpaces    = "\u0020\u0020\u0020\u0020"
)

type Node struct {
	text   string
	indent int
	parent *Node
	prev   *Node
	next   *Node
	first  *Node
	last   *Node
}

func New(s string) *Node {
	text, i := trimIndent(s)
	return &Node{
		text:   text,
		indent: i,
	}
}

func (n *Node) IsEmpty() bool {
	return n.text == ""
}

func (n *Node) Add(node *Node) *Node {
	if n.isChild(node) {
		n.addChild(node)
		return n
	}

	if n.parent == nil {
		return nil
	}

	return n.parent.Add(node)
}

func (n *Node) Text() string {
	return n.text
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) FirstChild() *Node {
	return n.first
}

func (n *Node) LastChild() *Node {
	return n.last
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) isChild(node *Node) bool {
	return node.indent > n.indent
}

func (n *Node) addChild(node *Node) {
	node.Remove()

	node.parent = n
	if n.last == nil {
		n.first = node
		n.last = node
	} else {
		node.prev = n.last
		n.last.next = node
		n.last = node
	}
}

func (n *Node) Remove() {
	p := n.parent
	if p != nil {
		if n == p.first {
			p.first = n.next
		}
		if n == p.last {
			n.last = n.prev
		}
	}
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	n.parent = nil
	n.prev = nil
	n.next = nil
}

func trimIndent(s string) (string, int) {
	i := 0
	return strings.TrimLeftFunc(s, func(r rune) bool {
		if r == ' ' {
			i++
			return true
		}
		return false
	}), i
}

func (n *Node) Print(w io.Writer) {
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
		c.Print(w)
	}
}
