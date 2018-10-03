// Go implementation of Shell Sort
// Author: Donald Shell (1959)

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 100)
	for i := 0; i <= 99; i++ {
		arr[i] = rand.Intn(100)
	}

	fmt.Printf("Initial array is: %#v\n\n", arr)

	for d := int(len(arr) / 2); d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}

	fmt.Println("Sorted result: ", arr)
}
