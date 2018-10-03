// Go implementation of Shell Sort
// Author: Donald Shell (1959)

package shell

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{46, 24, 33, 10, 2, 81, 50}
	Sort(arr)
	expected := []int{2, 10, 24, 33, 46, 50, 81}
	if !reflect.DeepEqual(arr, expected) {
		t.Fatalf("Failed sorting arr. %+v != %+v", arr, expected)
	}
}
