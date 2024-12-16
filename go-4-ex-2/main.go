package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	var a, b float64
	fmt.Print("Enter the length of side a: ")
	fmt.Scan(&a)
	fmt.Print("Enter the length of side b: ")
	fmt.Scan(&b)

	hypotenuse := computeHypotenuse(a, b)
	fmt.Printf("The hypotenuse of the triangle with sides %.2f and %.2f is %.2f\n", a, b, hypotenuse)
}
