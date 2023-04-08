package SearchForecastTree

import "fmt"

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
	for _, son := range n.children {
		sons = append(sons, son)
	}
	return sons
}

func getNodeSonsName(n *node) []rune {
	sons := make([]rune, 0)
	for son := range n.children {
		sons = append(sons, son)
	}
	return sons
}

func compareStringSlices(a, b []string) bool {
	// Handle cases where one or both input slices are nil or empty
	if (a == nil && len(b) == 0) || (b == nil && len(a) == 0) || (a == nil && b == nil) {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	visited := make([]bool, len(b))
	for _, strA := range a {
		found := false
		for j, strB := range b {
			if strA == strB && !visited[j] {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
