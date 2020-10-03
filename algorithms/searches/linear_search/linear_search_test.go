package linearsearch

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		desc     string
		tcArray  []int
		tcKey    int
		expected int
	}{
		{
			desc:     "One",
			tcArray:  []int{1, 2, 4, 6, 7, 8},
			tcKey:    7,
			expected: 4,
		},
		{
			desc:     "Two",
			tcArray:  []int{9, 5, 4, 6, 7, 8},
			tcKey:    3,
			expected: -1,
		},
		{
			desc:     "Three",
			tcArray:  []int{10, 9, 4, 6, 7, 8, 1, 3},
			tcKey:    20,
			expected: -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := LinearSearch(tC.tcArray, tC.tcKey)

			if tC.expected != actual {
				t.Fatalf("Expected that index %d, actual %d", tC.expected, actual)
			} else {
				fmt.Println("Success!")
			}
		})
	}
}
