package main

import "fmt"

func main() {
	// Call the function computeGrade
	var score float64
	fmt.Print("Enter the score (0-100): ")
	fmt.Scan(&score)

	grade := computeGrade(score)
	fmt.Printf("The grade for the score %.2f is: %s\n", score, grade)
}

// TODO: implement the function computeGrade
func computeGrade(score float64) string {
	if score < 0 || score > 100 {
		return "Invalid score"
	} else if score >= 90 {
		return "6.0"
	} else if score >= 80 {
		return "5.5"
	} else if score >= 70 {
		return "5.0"
	} else if score >= 60 {
		return "4.5"
	} else if score >= 50 {
		return "4.0"
	} else if score >= 40 {
		return "3.5"
	} else if score >= 30 {
		return "3.0"
	} else if score >= 20 {
		return "2.5"
	} else if score >= 10 {
		return "2.0"
	} else {
		return "1.0"
	}
}