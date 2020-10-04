package anagram

import "testing"

type testCase struct {
	s1       string
	s2       string
	expected bool
}

func Test_IsAnagram(t *testing.T) {
	testCases := []testCase{
		{"listen", "silent", true},
		{"foo", "bar", false},
		{"alert", "later", true},
		{"reset", "trees", true},
	}

	for _, tc := range testCases {
		actual := IsAnagram(tc.s1, tc.s2)
		if tc.expected != actual {
			t.Fatalf("Expected IsAnagram(\"%v\", %v) = %v, but got %v", tc.s1, tc.s2, tc.expected, actual)
		}
	}
}
