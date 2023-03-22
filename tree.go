package SearchForecastTree

func NewTree(sentences *[]string) *Tree {
	t := &Tree{
		root: newNode(0),
	}
	if sentences != nil {
		// insert default sentences
		for _, sentence := range *sentences {
			t.Insert(sentence)
		}
	}
	return t
}

func newNode(value rune) *node {
	return &node{
		value: value,
		sons:  make(map[rune]*node),
	}
}

func (t *Tree) Insert(sentence string) {
	cur := t.root
	for _, c := range sentence {
		if !cur.hasSon(c) {
			cur.sons[c] = newNode(c)
		}
		// move to next
		cur = cur.sons[c]
	}
}

func (n *node) hasSon(value rune) bool {
	_, ok := n.sons[value]
	return ok
}
