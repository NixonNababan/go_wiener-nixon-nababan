package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func getNthPrime(n, current, primeCount int) int {
	if primeCount == n {
		return current - 1
	}
	if isPrime(current) {
		return getNthPrime(n, current+1, primeCount+1)
	}
	return getNthPrime(n, current+1, primeCount)
}

func primeX(n int) int {
	return getNthPrime(n, 2, 0)
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
