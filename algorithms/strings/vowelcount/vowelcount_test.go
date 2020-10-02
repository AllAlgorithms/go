package vowelcount

import "testing"

type testCase struct {
	input    string
	expected int
}

func Test_Count(t *testing.T) {
	testCases := []testCase{
		{"abcdef", 2},
		{"hello, world!", 3},
		{"vwxyz", 0},
	}

	for _, tc := range testCases {
		actual := Count(tc.input)
		if tc.expected != actual {
			t.Fatalf("Expected Count(\"%v\") = %v, but got %v", tc.input, tc.expected, actual)
		}
	}
}
