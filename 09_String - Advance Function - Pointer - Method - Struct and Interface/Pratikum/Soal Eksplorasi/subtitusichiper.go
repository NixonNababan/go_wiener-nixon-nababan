package main

import (
	"fmt"
	"strings"
)

// Student adalah tipe data struct yang merepresentasikan seorang siswa
type Student struct {
	Name       string
	NameEncode string
}

// Cipher adalah interface yang merepresentasikan sebuah algoritma enkripsi/dekripsi
type Cipher interface {
	Encode(input string) string
	Decode(input string) string
}

// CaesarCipher adalah tipe data struct yang merepresentasikan algoritma Caesar Cipher
type CaesarCipher struct {
	Offset int // Offset adalah jumlah pergeseran dalam kode ASCII
}

// Encode adalah method untuk melakukan enkripsi menggunakan algoritma Caesar Cipher
func (c CaesarCipher) Encode(input string) string {
	var output strings.Builder

	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			// encrypt lowercase letters
			output.WriteRune('a' + ((char-'a')+c.Offset)%26)
		} else if char >= 'A' && char <= 'Z' {
			// encrypt uppercase letters
			output.WriteRune('A' + ((char-'A')+c.Offset)%26)
		} else {
			// do not encrypt other characters
			output.WriteRune(char)
		}
	}

	return output.String()
}

// Decode adalah method untuk melakukan dekripsi menggunakan algoritma Caesar Cipher
func (c CaesarCipher) Decode(input string) string {
	var output strings.Builder

	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			// decrypt lowercase letters
			output.WriteRune('a' + ((char-'a')-c.Offset+26)%26)
		} else if char >= 'A' && char <= 'Z' {
			// decrypt uppercase letters
			output.WriteRune('A' + ((char-'A')-c.Offset+26)%26)
		} else {
			// do not decrypt other characters
			output.WriteRune(char)
		}
	}

	return output.String()
}

func main() {
	var menu int
	var student Student

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&student.Name)

		// Buat instance dari CaesarCipher dengan offset sebesar 3
		cipher := CaesarCipher{Offset: 3}

		// Enkripsi nama siswa menggunakan CaesarCipher
		student.NameEncode = cipher.Encode(student.Name)

		fmt.Printf("\nEncode of student's name %s is : %s", student.Name, student.NameEncode)
	} else if menu == 2 {
		fmt.Print("\nInput Encrypted Student Name: ")
		fmt.Scan(&student.NameEncode)

		// Buat instance dari CaesarCipher dengan offset sebesar 3
		cipher := CaesarCipher{Offset: 3}

		// Dekripsi nama siswa menggunakan CaesarCipher
		student.Name = cipher.Decode(student.NameEncode)

		fmt.Printf("\nDecode of student's name %s is : %s", student
