package SearchForecastTree

func NewTree(sentences []string) *Tree {
	t := &Tree{
		root: newNode(0),
	}
	// insert default sentences
	for _, sentence := range sentences {
		t.Insert(sentence)
	}
	return t
}

func newNode(value rune) *node {
	return &node{
		value:    value,
		children: make(map[rune]*node),
	}
}

func (t *Tree) Insert(sentence string) {
	cur := t.root
	for _, c := range sentence {
		if !cur.hasSon(c) {
			cur.children[c] = newNode(c)
		}
		// move to next
		cur = cur.children[c]
	}
	cur.end = true
}

func (n *node) hasSon(value rune) bool {
	_, ok := n.children[value]
	return ok
}

func (n *node) getSonCount() int {
	return len(n.children)
}

// Search from the tree by this prefix, return results which are matched
// TODO: limit count, priority
func (t *Tree) Search(prefix string) []string {
	cur := t.matchForCur(prefix)
	if cur == nil {
		return nil
	}
	// we can match this prefix, search for results
	// eg: prefix = "test" result = ["test11", "test233", "test"]
	return t.searchWithCur(prefix, cur)
}

func (t *Tree) searchWithCur(prefix string, cur *node) []string {
	var result []string
	if cur.end {
		result = append(result, prefix)
	}
	if cur.getSonCount() == 0 {
		return result
	}
	for c, n := range cur.children {
		result = append(result, t.searchWithCur(prefix+string(c), n)...)
	}
	return result
}

// Get cur by the match string, return nil if not be here
func (t *Tree) matchForCur(prefix string) *node {
	cur := t.root
	for _, c := range prefix {
		if !cur.hasSon(c) {
			// match null
			return nil
		}
		cur = cur.children[c]
	}
	return cur
}
