package SearchForecastTree

type Tree struct {
	root *node
}

type node struct {
	value rune
	sons  map[rune]*node
	//priority float64
}
