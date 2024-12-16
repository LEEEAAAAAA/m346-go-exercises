package main

import (
	"fmt"
	"math"
)

func computeQuadraticFormula(a, b, c float64) (float64, float64, bool) {
	discriminant := math.Pow(b, 2) - 4*a*c
	if discriminant < 0 {
		// No real roots
		return 0, 0, false
	}

	root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
	return root1, root2, true
}

func main() {

	var a, b, c float64
	fmt.Print("Enter coefficient a: ")
	fmt.Scan(&a)
	fmt.Print("Enter coefficient b: ")
	fmt.Scan(&b)
	fmt.Print("Enter coefficient c: ")
	fmt.Scan(&c)

	root1, root2, hasRealRoots := computeQuadraticFormula(a, b, c)
	if !hasRealRoots {
		fmt.Println("The equation has no real roots.")
	} else {
		fmt.Printf("The roots of the equation are: %.2f and %.2f\n", root1, root2)
	}
}
