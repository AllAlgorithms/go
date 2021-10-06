package leapyear

import "testing"

type testCase struct {
	input    int
	expected bool
}

func Test_IsLeapYear(t *testing.T) {
	testCases := []testCase{
		{2000, true},
		{1900, false},
		{2100, false},
		{1700, false},
		{1891, false},
	}

	for _, testCase := range testCases {
		result := IsLeapYear(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %v, got %v", testCase.expected, result)
		}
	}
}
