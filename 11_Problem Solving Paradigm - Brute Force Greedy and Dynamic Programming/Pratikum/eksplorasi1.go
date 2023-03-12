package main

import "fmt"

func main() {
	var number int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&number)
	fmt.Printf("Konversi Romawi: %s\n", toRoman(number))
}

// fungsi untuk mengkonversi angka normal menjadi angka romawi
func toRoman(number int) string {
	// membuat slice untuk menampung angka-angka romawi dan arabicnya
	romanNumerals := []struct {
		arabic int
		roman  string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result string

	// mengonversi angka normal ke angka romawi dengan cara mengurangi nilai angka normal dengan nilai arabic
	// pada setiap angka romawi yang ada
	for _, value := range romanNumerals {
		for number >= value.arabic {
			result += value.roman
			number -= value.arabic
		}
	}

	return result
}
