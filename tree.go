package SearchForecastTree

func NewTree(sentences *[]string) *Tree {
	t := &Tree{
		root: newNode(0, 0),
	}
	if sentences != nil {
		// insert default sentences
		for _, sentence := range *sentences {
			t.Insert(sentence)
		}
	}
	return t
}

func newNode(value rune, priority float64) *node {
	return &node{
		value:    value,
		sons:     make(map[rune]*node),
		priority: priority,
	}
}

func (t *Tree) Insert(sentence string) {
	cur := t.root
	for _, c := range sentence {
		if !cur.hasSon(c) {

		}
	}
}

func (n *node) hasSon(value rune) bool {
	_, ok := n.sons[value]
	return ok
}
