package main

import "fmt"

func PairSum(arr []int, target int) []int {
	// buat sebuah map untuk menyimpan nilai-nilai yang telah diproses
	m := make(map[int]int)

	// iterasi array
	for i, num := range arr {
		// hitung selisih target dengan angka saat ini
		diff := target - num

		// cek apakah selisih tersebut sudah ada dalam map
		if j, ok := m[diff]; ok {
			// jika ada, berarti ditemukan pasangan yang sesuai
			return []int{j, i}
		}

		// jika tidak, tambahkan nilai saat ini ke dalam map
		m[num] = i
	}

	// jika tidak ada pasangan yang ditemukan, kembalikan slice kosong
	return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // Output : [1, 3]

	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // Output : [0, 2]

	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // Output : [2, 3]

	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) // Output : [1, 2]

	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) // Output : [0, 1]
}
