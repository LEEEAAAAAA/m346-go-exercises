package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	eyesFiles, _ := os.Create("eyes.txt")
	defer eyesFiles.Close()


	dice, _ := os.Create("dice.log")
	defer dice.Close()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")
	fmt.Fprintln(os.Stdout, "the dice was rolled at", when)

	// go run ex3/main.go TODO
}
