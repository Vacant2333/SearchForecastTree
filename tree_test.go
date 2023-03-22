package SearchForecastTree

import (
	"fmt"
	"testing"
)

func TestNewTree(t *testing.T) {
	sentences := &[]string{
		"abo",
		"abd",
		"ab,s",
		"cbd",
	}
	tree := NewTree(sentences)
	printTree(tree)
}

func printTree(t *Tree) {
	printNode(t.root)
}

func printNode(n *node) {
	sons := getNodeSonsName(n)
	fmt.Printf("Node:[%c] Count:[%v] Sons:[%v]\n", n.value, len(sons), string(sons))
	for _, son := range getNodeSons(n) {
		printNode(son)
	}
}

func getNodeSons(n *node) []*node {
	sons := make([]*node, 0)
	for _, son := range n.sons {
		sons = append(sons, son)
	}
	return sons
}

func getNodeSonsName(n *node) []rune {
	sons := make([]rune, 0)
	for son := range n.sons {
		sons = append(sons, son)
	}
	return sons
}
