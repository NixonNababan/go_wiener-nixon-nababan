package main

import "fmt"

func caesar(offset int, input string) string {
	result := ""
	for _, char := range input {
		if char == ' ' {
			result += " "
			continue
		}

		// Menghitung kode ASCII untuk karakter saat ini
		charCode := int(char)

		// Menentukan kode ASCII untuk huruf awal
		base := 97

		// Menghitung jumlah pergeseran dalam kode ASCII
		shift := (charCode - base + offset) % 26

		// Menambahkan pergeseran pada kode ASCII huruf awal
		newCharCode := base + shift

		// Mengonversi kode ASCII kembali ke karakter
		newChar := string(newCharCode)

		result += newChar
	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
