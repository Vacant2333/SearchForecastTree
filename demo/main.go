package main

import (
	"SearchForecaseTree"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sentences := SearchForecastTree.GetSentencesFromFile("titles1.txt")
	tree := SearchForecastTree.NewTree(sentences)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Tree size: %.2fMb\n", tree.MemorySizeMB())
	for {
		text, _ := reader.ReadString('\n')
		fmt.Println(strings.Join(tree.Search(text[:len(text)-1]), ";"))
	}
}
