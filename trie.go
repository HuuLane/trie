package trie

import (
	"strings"
)

type Node struct {
	char     rune
	children Nodes
	isEnd    bool
}

type Nodes []*Node

func NewNode(char rune) *Node {
	n := new(Node)
	n.char = char
	n.children = make(Nodes, 26)
	return n
}

func NewTrie() *Node {
	return NewNode('/')
}

func (n *Node) Insert(s string) {
	for _, char := range s {
		i := char - 'a'
		if n.children[i] == nil {
			n.children[i] = NewNode(char)
		}
		n = n.children[i]
	}
	n.isEnd = true
}

func (n *Node) Exists(s string) bool {
	for _, char := range s {
		i := char - 'a'
		if n.children[i] == nil {
			return false
		} else {
			n = n.children[i]
		}
	}
	return n.isEnd
}

func filterNilNode(nodes Nodes) Nodes {
	res := make(Nodes, 0)
	for _, node := range nodes {
		if node != nil {
			res = append(res, node)
		}
	}
	return res
}

func (n *Node) allNodesChar() [][]rune {
	var layers [][]rune

	// root generation
	layers = append(layers, []rune{n.char})
	generation := filterNilNode(n.children)

	for len(generation) != 0 {
		layer := make([]rune, 0)
		nextGeneration := make(Nodes, 0)
		for _, n := range generation {
			if n == nil {
				continue
			}
			layer = append(layer, n.char)
			nextGeneration = append(nextGeneration,
				filterNilNode(n.children)...)
		}
		layers = append(layers, layer)
		generation = nextGeneration
	}
	return layers
}

func (n *Node) String() string {
	layers := n.allNodesChar()
	var b strings.Builder
	for _, layer := range layers {
		b.WriteString(string(layer) + "\n")
	}
	return b.String()
}
