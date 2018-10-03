// Package linearsort is a Go implementation
// of linear sorting algorith
//
// Author: Filip Marek
package linearsort

// Sort function sorts given int slice
func Sort(slice []int) []int {
	for {
		iterationSorted := false

		for index, count := 0, len(slice)-1; index < count; index++ {
			if slice[index] > slice[index+1] {
				slice[index], slice[index+1] = slice[index+1], slice[index]

				iterationSorted = true
			}
		}

		if !iterationSorted {
			break
		}
	}

	return slice
}
