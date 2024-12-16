package main

import "fmt"

	// TODO: declare a type for Student (with first and last name)
type Student struct {
	FirstName string
	LastName  string
}

	// TODO: declare a type for Class (consisting of multiple students)
type Class []Student

func main() {
	// TODO: declare a map of modules being attended by multiple classes
	modules := map[string]Class{
		"Mathematics": {
			{FirstName: "Alice", LastName: "Johnson"},
			{FirstName: "Bob", LastName: "Smith"},
		},
		"Science": {
			{FirstName: "Charlie", LastName: "Brown"},
			{FirstName: "Diana", LastName: "White"},
		},
		"History": {
			{FirstName: "Eve", LastName: "Green"},
			{FirstName: "Frank", LastName: "Black"},
		},
	}

	// TODO: output everything using fmt.Println()
	for module, class := range modules {
		fmt.Println("Module:", module)
		for _, student := range class {
			fmt.Printf("  Student: %s %s\n", student.FirstName, student.LastName);
		}
	}
}