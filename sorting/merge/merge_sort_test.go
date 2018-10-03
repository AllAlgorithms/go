package merge

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{12, 11, 13, 5, 6, 7}
	expected := []int{5, 6, 7, 11, 12, 13}
	if got := Sort(arr); !reflect.DeepEqual(got, expected) {
		t.Errorf("Failed sorting arr. %+v != %+v", got, expected)
	}
}
