package main

import "fmt"

// Fungsi ArrayMerge menerima dua parameter slice, arrayA dan arrayB, dan kemudian membuat slice baru (mergedArray) dengan ukuran arrayA + arrayB menggunakan fungsi make. Kemudian, elemen arrayA disalin ke mergedArray menggunakan fungsi copy.
func ArrayMerge(arrayA, arrayB []string) []string {
	mergedArray := make([]string, len(arrayA)+len(arrayB))
	copy(mergedArray, arrayA)

	// melakukan loop pada setiap elemen arrayB dan memeriksa apakah elemen tersebut sudah ada di dalam arrayA. Jika elemen tidak ditemukan dalam arrayA, maka elemen tersebut ditambahkan ke mergedArray menggunakan fungsi append.
	for _, value := range arrayB {
		found := false
		for _, element := range arrayA {
			if element == value {
				found = true
				break
			}
		}
		if !found {
			mergedArray = append(mergedArray, value)
		}
	}

	return mergedArray
}

func main() {

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))

}
