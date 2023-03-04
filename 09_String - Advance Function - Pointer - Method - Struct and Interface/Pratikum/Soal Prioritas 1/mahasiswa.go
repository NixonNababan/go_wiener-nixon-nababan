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
