package main

import "fmt"

func fibonacci() func() int64 {
	var prev int64 = -1
	var res int64 = 1
	return func() int64 {
		res, prev = res+prev, res
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}
