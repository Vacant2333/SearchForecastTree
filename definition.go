package SearchForecastTree

type Tree struct {
	root *node
}

type node struct {
	value    rune
	children map[rune]*node
	// if it's true, means it's a sentence
	end bool
	//priority float64
}
