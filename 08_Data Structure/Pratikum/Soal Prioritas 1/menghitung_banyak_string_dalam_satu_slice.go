package main

import "fmt"

//  setiap string pada slice akan diloop dan dicocokkan dengan setiap string pada slice. Jika string yang sedang dicari ditemukan, maka variabel count akan diincrement
func Mapping(slice []string) map[string]int {
	result := make(map[string]int)
	for _, s := range slice {
		count := 0
		for _, r := range slice {
			if s == r {
				count++
			}
		}
		// Setelah selesai loop kedua, variabel count akan menjadi jumlah kemunculan string tersebut pada slice, dan dimasukkan ke dalam map yang akan dikembalikan oleh fungsi Mapping.
		result[s] = count
	}
	return result
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))

	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))

	fmt.Println(Mapping([]string{}))
}
