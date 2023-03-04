# (9) String - Advance Function - Pointer - Method - Struct and Interface
[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

## Prioritas 1

### Soal 1
```Go
package main

import "fmt"

// definisi struct Mobil
type Mobill struct {
	Type   string
	FuelIn float64
}

// method untuk menghitung perkiraan jarak yang bisa ditempuh
func (c *Mobill) EstimateDistance() float64 {
	fuelEfficiency := 1.5 // km per mill
	distance := fuelEfficiency * c.FuelIn
	return distance
}

func main() {
	// membuat objek mobil dengan tipe "sedan" dan bahan bakar 1000 milliliter
	Mobilku := Mobill{"sedan", 1000.0}

	// menghitung perkiraan jarak yang bisa ditempuh oleh mobil menggunakan method EstimateDistance()
	distance := Mobilku.EstimateDistance()

	// mencetak hasil perhitungan ke layar
	fmt.Println("Perkiraan jarak yang bisa ditempuh:", distance, "km")
}

```

### Soal 2
```Go
package main

import (
	"fmt"
)

type Student struct {
	name  string
	score []int
}

func main() {
	students := make([]Student, 5)
	var totalScore, minScore, maxScore int
	var minStudent, maxStudent Student

	// Input data siswa
	for i := 0; i < 5; i++ {
		fmt.Printf("Input %d Student's Name: ", i+1)
		fmt.Scan(&students[i].name)

		fmt.Printf("Input %d Student's Score: ", i+1)
		var score int
		fmt.Scan(&score)
		students[i].score = append(students[i].score, score)

		// Menghitung total score
		totalScore += score
	}

	// Menghitung rata-rata score
	averageScore := float64(totalScore) / float64(len(students))
	fmt.Printf("Average Score : %.2f\n", averageScore)

	// Mencari nilai minimum dan siswa dengan nilai minimum
	minScore = students[0].score[0]
	for _, student := range students {
		for _, score := range student.score {
			if score < minScore {
				minScore = score
				minStudent = student
			}
		}
	}
	fmt.Printf("Min Score of Students : %s (%d)\n", minStudent.name, minScore)

	// Mencari nilai maksimum dan siswa dengan nilai maksimum
	maxScore = students[0].score[0]
	for _, student := range students {
		for _, score := range student.score {
			if score > maxScore {
				maxScore = score
				maxStudent = student
			}
		}
	}
	fmt.Printf("Max Score of Students : %s (%d)\n", maxStudent.name, maxScore)
}
```

### Soal 3
```Go
package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	//fungsi bernama Compare yang memiliki dua parameter bertipe string, yaitu a dan b, dan mengembalikan sebuah nilai bertipe string.
	if strings.Contains(a, b) {
		return b
	} else if strings.Contains(b, a) {
		return a
	} else {
		return ""
	}
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}

```

### Soal 4
```Go
package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	// inisialisasi nilai awal min dan max
	min = *numbers[0]
	max = *numbers[0]

	// loop melalui array, membandingkan nilai setiap elemen dengan min dan max
	for _, num := range numbers {
		if *num < min {
			min = *num
		}
		if *num > max {
			max = *num
		}
	}

	return
}
func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	// pass alamat variabel ke fungsi getMinMax menggunakan operator &
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Printf("%d is min\n", min)
	fmt.Printf("%d is max\n", max)
}
```

## Soal Prioritas 2

### Soal 1
```Go
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
```

## Soal Explorasi
### Soal 1
```Go
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

```