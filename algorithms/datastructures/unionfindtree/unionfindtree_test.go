package unionfindtree

import (
	"testing"
)

func TestUnionFindTree(t *testing.T) {
	t.Run("Test UnionFindTree", func(t *testing.T) {
		uft := Create(10)
		/*
			Grouping Odd and Even numbers.
		*/

		uft.Merge(0, 2)
		uft.Merge(2, 4)
		uft.Merge(4, 6)
		uft.Merge(6, 8)

		uft.Merge(1, 3)
		uft.Merge(3, 5)
		uft.Merge(5, 7)
		uft.Merge(7, 9)

		for i := 0; i < 8; i += 2 {
			if !uft.IsSame(i, i+2) {
				t.Errorf("Invalid Grouping: %d, %d", i, i+2)
			}
			if !uft.IsSame(i+1, i+3) {
				t.Errorf("Invalid Grouping: %d, %d", i+1, i+3)
			}
			if uft.IsSame(i, i+1) {
				t.Errorf("Invalid Grouping: %d, %d", i, i+1)
			}
		}
	})
}
