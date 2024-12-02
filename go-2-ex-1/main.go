package main

import "fmt"

type FullName struct {
	firstName string
	lastName string
}

// TODO: declare a structure for birth date
type Birthdate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	// TODO: embed full name and birth date information
	FullName
	Birthdate
	NumberOfSiblings byte
	ZodiacSign       string
}

func main() {
	var me = Profile{
		FullName: FullName{
			firstName: "John",
			lastName:  "Doe",
		},
		Birthdate: Birthdate{
			Day:   15,
			Month: 6,
			Year:  1990,
		},
		NumberOfSiblings: 0,   // TODO: adjust
		ZodiacSign:       "Canser", // TODO: adjust
	}
	

	fmt.Println(me)
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
