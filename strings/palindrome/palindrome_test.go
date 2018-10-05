package palindrome

import "testing"

type testCase struct {
	input    string
	expected bool
}

func Test_IsPalindrome(t *testing.T) {
	testCases := []testCase{
		{"racecar", true},
		{"ferrari", false},
		{"civic", true},
	}

	for _, tc := range testCases {
		actual := IsPalindrome(tc.input)
		if tc.expected != actual {
			t.Fatalf("Expected IsPalindrome(\"%v\") = %v, but got %v", tc.input, tc.expected, actual)
		}
	}
}
