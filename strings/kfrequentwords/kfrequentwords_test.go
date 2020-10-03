package kfrequentwords

import (
	"reflect"
	"testing"
)

type testCase struct {
	s        string
	k        int
	expected []string
}

func Test_Top(t *testing.T) {
	testCases := []testCase{
		{"i love algorithm and i love coding", 2, []string{"i", "love"}},
		{"the day is sunny the the the sunny is is", 4, []string{"the", "is", "sunny", "day"}},
	}

	for _, tc := range testCases {
		actual := Top(tc.s, tc.k)
		if !reflect.DeepEqual(tc.expected, actual) {
			t.Fatalf("Expected Top(\"%v\", %v) = %v, but got %v", tc.s, tc.k, tc.expected, actual)
		}
	}
}
