package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	if n == 1 {
		return 0
	}
	if n == 2 {
		return int(math.Abs(float64(jumps[0] - jumps[1])))
	}
	return int(math.Min(float64(Frog(jumps[:n-1])+int(math.Abs(float64(jumps[n-1]-jumps[n-2])))), float64(Frog(jumps[:n-2])+int(math.Abs(float64(jumps[n-1]-jumps[n-3]))))))
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
