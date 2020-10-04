package breadthfirstsearch

import (
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := map[int]bool{
		-1:  true,
		2:   true,
		20:  false,
		5:   true,
		15:  true,
		7:   true,
		13:  false,
		18:  false,
		-22: false,
		3:   true,
		19:  true,
		21:  false,
	}

	root := &Node{Value: 0}
	for value, exists := range testCases {
		if exists {
			root.Insert(value)
		}
	}

	for value, exists := range testCases {
		actual := root.Search(value)
		if exists != actual {
			t.Errorf("Expected searching result for value %d would be %t, got %t", value, exists, actual)
		}
	}
}
