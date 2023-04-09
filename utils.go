package SearchForecastTree

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unsafe"
)

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

func GetSentencesFromFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	var result []string
	// Create a new Scanner to read the file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

func CalculateTreeSize(n *node) uintptr {
	if n == nil {
		return 0
	}
	// Calculate the memory size of the current node
	nodeSize := unsafe.Sizeof(*n)
	// Calculate the memory size of the node's 'sons' map
	mapSize := unsafe.Sizeof(map[rune]*node{}) + uintptr(len(n.children))*(unsafe.Sizeof(rune(0))+unsafe.Sizeof(&node{}))
	// Recursively calculate the memory size of each child node in the 'sons' map
	var childrenSize uintptr
	for _, child := range n.children {
		childrenSize += CalculateTreeSize(child)
	}
	// Return the total memory size
	return nodeSize + mapSize + childrenSize
}

func (t *Tree) MemorySizeMB() float64 {
	bytes := float64(CalculateTreeSize(t.root))
	megabytes := bytes / (1024 * 1024)
	return megabytes
}
