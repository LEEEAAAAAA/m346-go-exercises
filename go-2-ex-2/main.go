package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	fibs[2] = fibs[0] + fibs[1]
	fibs[3] = fibs[1] + fibs[2] // 3
	fibs[4] = fibs[2] + fibs[3] // 5

	// TODO: correct up to index 4 using direct element access
	fibs = append(fibs, fibs[3]+fibs[4]) // TODO: replace 0 with the next Fibonacci number

	// TODO: compute three more Fibonacci numbers and append them
	fibs = append(fibs, fibs[4]+fibs[5]) // 13
	fibs = append(fibs, fibs[5]+fibs[6]) // 21
	fibs = append(fibs, fibs[6]+fibs[7]) // 34

	fmt.Println(fibs)
}
