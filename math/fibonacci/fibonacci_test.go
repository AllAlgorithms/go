package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	f := Fibonacci()

	var res []int64
	for i := 0; i < 10; i++ {
		res = append(res, f())
	}

	expected := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("failed to assert that %+v is equal to the expected value %+v", res, expected)
	}
}
