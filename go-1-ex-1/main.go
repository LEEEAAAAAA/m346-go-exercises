package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	var firstName string = "Lea"
	var lastName string = "Schürch"
	var dayOfBirth int = 24
	var monthOfBirth int = 6
	var yearOfBirth int = 2007
	var numberOfSiblings int = 1
	var heightInMeters float64 = 2.0
	var zodiacSign string = "Cancer" 

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %s\n", zodiacSign) 
}

