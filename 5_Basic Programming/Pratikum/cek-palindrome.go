package main

import "fmt"

func main() {
	var s string
	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&s)

	if isPalindrome(s) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan palindrome")
	}
}

func isPalindrome(s string) bool {
	for b := 0; b < len(s)/2; b++ {
		if s[b] != s[len(s)-b-1] {
			return false
		}
	}
	return true
}
