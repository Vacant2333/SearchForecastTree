# SearchForecastTree
SearchForecastTree(TrieTree)

## How to use

```go
package main

import "fmt"

func main() {
	sentences := []string{"ab", "abc", "a", "ab,s", "d"}
	tree := SearchForecastTree.NewTree(sentences)
	fmt.Println(tree.Search("ab"))
    // output: [ab abc ab,s]
}
```
you can run demo/main.go to know how to use.
