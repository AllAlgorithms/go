package bucketsort

import "testing"

func TestSort(t *testing.T) {
	testVelue := []int{46, 24, 33, 10, 2, 81, 50}
	expected := []int{2, 10, 24, 33, 46, 50, 81}
	actual := Sort(testVelue)

	if len(testVelue) != len(actual) {
		t.Fatalf("Expected that output array length %d, actual %d", len(expected), len(actual))
	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Fatalf("\nExpected sorted array: %v\nActual sorted array:   %v", expected, actual)
		}
	}
}
