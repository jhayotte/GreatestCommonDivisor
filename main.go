package main

import (
	"fmt"
)

func gcd(x, y int) int {
	var r int

	for y != 0 {
		/*Formula on a single line*/
		//x, y = y, x%y
		r = x % y
		x = y
		y = r
	}
	return x
}

func main() {
	var n = []int{48, 96, 264, 96}
	n = append(n, 24)

	var result = n[0]
	for index := 0; index < len(n); index++ {
		result = gcd(result, n[index])
	}
	fmt.Printf("GCD: %d", result)
}
