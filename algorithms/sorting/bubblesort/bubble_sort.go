// Go implementation of bubble sort
//
// Author: Carlos Abraham Hernandez

package bubblesort

func Sort(arr []int) []int{
	if len(arr) <= 1 {
		return arr
	}

	n := len(arr) - 1

	for {
		if n == 0 {
			break
		}

		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		n -= 1
	}

	return arr
}
