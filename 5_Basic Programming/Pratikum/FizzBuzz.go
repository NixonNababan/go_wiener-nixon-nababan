package main

import (
	"fmt"
)

func main() {
	for x := 1; x <= 100; x++ {
		if x%3 == 0 && x%5 == 0 {
			fmt.Println("FizzBuzz")
		}
		switch {
		case x%3 == 0:
			fmt.Println("Fizz")
		case x%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(x)
		}
	}
}
