package cyclesort

import (
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		desc     string
		testCase []int
		expected []int
	}{
		{
			desc:     "One",
			testCase: []int{11, 4, 434, 2, 14, 3, 300, 1, 435, 5, 6, 20, 1, 3},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 6, 11, 14, 20, 300, 434, 435},
		},
		{
			desc:     "Two",
			testCase: []int{46, 24, 33, 10, 2, 81, 50},
			expected: []int{2, 10, 24, 33, 46, 50, 81},
		},
		{
			desc:     "Three",
			testCase: []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			desc:     "Four",
			testCase: []int{1},
			expected: []int{1},
		},
		{
			desc:     "Five",
			testCase: []int{2, 1},
			expected: []int{1, 2},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Sort(tC.testCase)

			if len(tC.expected) != len(actual) {
				t.Fatalf("Expected that output array length %d, actual %d", len(tC.expected), len(actual))
			}

			for i := range tC.expected {
				if tC.expected[i] != actual[i] {
					t.Fatalf("\nExpected sorted array: %v\nActual sorted array:   %v", tC.expected, actual)
				}
			}
		})
	}
}
