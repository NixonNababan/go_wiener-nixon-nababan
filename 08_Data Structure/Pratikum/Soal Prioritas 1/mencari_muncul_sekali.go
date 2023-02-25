package main

import (
	"fmt"
	"strconv"
)

// Pada setiap perulangan, kita menghitung berapa kali karakter pada indeks i muncul pada string input angka.
func munculSekali(angka string) []int {
	var result []int
	for i := 0; i < len(angka); i++ {
		count := 0
		for j := 0; j < len(angka); j++ {
			if angka[i] == angka[j] {
				count++
			}
		}

		//angka akan di-convert dari string menjadi integer menggunakan fungsi Atoi dari package strconv, lalu ditambahkan ke dalam slice result menggunakan fungsi append.
		if count == 1 {
			num, _ := strconv.Atoi(string(angka[i]))
			result = append(result, num)
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // Output :[4]
	fmt.Println(munculSekali("76523752"))   // Output :[6 3]
	fmt.Println(munculSekali("12345"))      // Output :[1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // Output :[]
	fmt.Println(munculSekali("0872504"))    // Output :[8 7 2 5 4]
}
