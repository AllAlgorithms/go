package main

import "fmt"

func gnomeSort(arr []int) []int {
	index := 0
	for index < len(arr) {
		if index == 0 {
			index++
		} 
		if arr[index] >= arr[index - 1] {
			index++
		} else {
			arr[index], arr[index-1] = arr[index-1], arr[index]
			index--		
		}
	}
	return arr
}

func main() {
	var arr = []int {123, 23, 42, 12, 64, 1213}
	fmt.Println(gnomeSort(arr))
}
