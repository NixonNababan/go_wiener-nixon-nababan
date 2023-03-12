package main

import "fmt"

func simpleEquation(a, b, c int) (int, int, int) {
	x := (a + b + c) / 3
	y := (a * b * c) / x
	z := (a*a + b*b + c*c) / (3 * x * x)
	return x, y, z
}
func main() {
	fmt.Println(simpleEquation(123, 0, 0)) // No solution.
	fmt.Println(simpleEquation(6, 6, 14))  // 123
}
