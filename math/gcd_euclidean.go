package main

import "fmt"

// greatest common divisor (GCD) via Euclidean algorithm
func gcd_euclidean(a, b uint) uint {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	fmt.Println("gcd(100,200) =", gcd_euclidean(100, 200))
	fmt.Println("gcd(4,2) =", gcd_euclidean(4, 2))
	fmt.Println("gcd(6,3) =", gcd_euclidean(6, 3))
}
