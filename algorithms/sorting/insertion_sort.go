package main

/*
 * Insertion sort - https://en.wikipedia.org/wiki/Insertion_sort
 */

import "fmt"

func main() {

	arr := []int{46, 24, 33, 10, 2, 81, 50}
	fmt.Println("Unsorted array")
	fmt.Println(arr)

	if len(arr) <= 1 {
		fmt.Println("Sorted array")
		fmt.Println(arr)
		return
	}

	var i, j int
	for i = 1; i < len(arr); i++ {
		for j = 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println("Sorted array")
	fmt.Println(arr)
}
