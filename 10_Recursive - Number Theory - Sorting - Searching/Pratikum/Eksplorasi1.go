package main

import (
	"fmt"
)

func MaxSequence(arr []int) int {
	maximum, TheEndofmax := 0, 0
	for _, num := range arr {

		TheEndofmax += num

		if TheEndofmax < 0 {

			TheEndofmax = 0
		}

		if maximum < TheEndofmax {

			maximum = TheEndofmax

		}
	}
	return maximum
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 13
}
