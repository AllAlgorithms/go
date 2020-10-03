package morse

import (
	"fmt"
	"testing"
)

type testCase struct {
	in       string
	expected string
	err      error
}

func Test_Encode(t *testing.T) {
	testCases := []testCase{
		{"TEST", "- . ... -", nil},
		{"abc123xyz987", ".- -... -.-. .---- ..--- ...-- -..- -.-- --.. ----. ---.. --...", nil},
		{"@", "", fmt.Errorf("unknown char @")},
	}

	for _, tc := range testCases {
		actual, err := Encode(tc.in)
		if tc.expected != actual {
			t.Fatalf("Expected Encode(\"%v\") = %v, but got %v", tc.in, tc.expected, actual)
		}
		if tc.err != nil && err == nil {
			t.Fatalf("Expected Encode(\"%v\") to return error, but error is nil", tc.in)
		}
	}
}

func Test_Decode(t *testing.T) {
	testCases := []testCase{
		{"- . ... -", "test", nil},
		{".- -... -.-. .---- ..--- ...-- -..- -.-- --.. ----. ---.. --...", "abc123xyz987", nil},
		{"-.....-", "", fmt.Errorf("unknown code -.....-")},
	}

	for _, tc := range testCases {
		actual, err := Decode(tc.in)
		if tc.expected != actual {
			t.Fatalf("Expected Encode(\"%v\") = %v, but got %v", tc.in, tc.expected, actual)
		}
		if tc.err != nil && err == nil {
			t.Fatalf("Expected Encode(\"%v\") to return error, but error is nil", tc.in)
		}
	}
}
