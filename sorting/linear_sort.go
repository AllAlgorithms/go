// Go implementation of linear sort
//
// Author: Filip Marek

package main

import "fmt"

func main() {
	arr := []int{46, 24, 33, 10, 2, 81, 50}

	fmt.Println("Unsorted Array")
	fmt.Println(arr)

	for {
		iterationSorted := false

		for index, count := 0, len(arr)-1; index < count; index++ {
			if arr[index] > arr[index+1] {
				arr[index], arr[index+1] = arr[index+1], arr[index]

				iterationSorted = true
			}
		}

		if !iterationSorted {
			break
		}
	}

	fmt.Println("Sorted Array")
	fmt.Println(arr)
}
