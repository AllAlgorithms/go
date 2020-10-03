package binarysearch

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
			tcArray:  []int{1, 2, 4, 6, 7, 8},
			tcKey:    3,
			expected: -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := BinarySearch(tC.tcArray, 0, len(tC.tcArray), tC.tcKey)

			if tC.expected != actual {
				t.Fatalf("Expected that index %d, actual %d", tC.expected, actual)
			} else {
				fmt.Println("Success!")
			}
		})
	}
}
