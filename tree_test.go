package SearchForecastTree

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	// Test case 1: Empty input
	tree1 := NewTree(nil)
	if tree1 == nil || tree1.root == nil {
		t.Error("NewTree should return a tree with root initialized")
	}

	// Test case 2: Non-empty input
	testSentences := []string{"hello", "world", "search", "tree"}
	tree2 := NewTree(testSentences)

	// Check if tree2 is not nil
	if tree2 == nil || tree2.root == nil {
		t.Error("NewTree should return a tree with root initialized")
	}

	// Check if tree2 contains the given sentences
	for _, sentence := range testSentences {
		cur := tree2.matchForCur(sentence)
		if cur == nil {
			t.Errorf("NewTree should insert the given sentences into the tree: %s not found", sentence)
		}
	}

	// Test case 3: Check for non-existing sentence
	nonExistingSentence := "notfound"
	cur := tree2.matchForCur(nonExistingSentence)
	if cur != nil {
		t.Errorf("NewTree should not contain the non-existing sentence: %s", nonExistingSentence)
	}
}

func TestSearch(t *testing.T) {
	testCases := []struct {
		name        string
		sentences   []string
		prefix      string
		expectedRes []string
	}{
		{
			name:        "Search with empty tree",
			sentences:   []string{},
			prefix:      "",
			expectedRes: nil,
		},
		{
			name:        "Search with existing prefix",
			sentences:   []string{"hello", "world", "search", "tree"},
			prefix:      "t",
			expectedRes: []string{"tree"},
		},
		{
			name:        "Search with non-existing prefix",
			sentences:   []string{"hello", "world", "search", "tree"},
			prefix:      "notfound",
			expectedRes: nil,
		},
		{
			name:        "Search with empty prefix",
			sentences:   []string{"hello", "world", "search", "tree"},
			prefix:      "",
			expectedRes: []string{"hello", "world", "search", "tree"},
		},
		{
			name:        "Search with common prefix",
			sentences:   []string{"apple", "app", "apricot", "banana"},
			prefix:      "app",
			expectedRes: []string{"apple", "app"},
		},
		{
			name:        "Search with chinese",
			sentences:   []string{"测试123", "测试title", "测t", "banana"},
			prefix:      "测试",
			expectedRes: []string{"测试123", "测试title"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tree := NewTree(tc.sentences)
			result := tree.Search(tc.prefix)
			if !compareStringSlices(result, tc.expectedRes) {
				t.Errorf("Expected %v, but got %v", tc.expectedRes, result)
			}
		})
	}
}
