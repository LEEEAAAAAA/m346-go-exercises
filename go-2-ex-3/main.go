package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
		modules := map[int]string{
			104: "Introduction to Programming",
			117: "Data Structures",
			346: "Algorithms",
		}
	
		// Access elements
		fmt.Println("Modul 104:", modules[104])
		fmt.Println("Modul 117:", modules[117])
		fmt.Println("Modul 346:", modules[346])
	
	// TODO: delete one
		delete(modules, 117)
	
	// TODO: add one
		modules[220] = "Operating Systems"
	
	// TODO: replace one
		modules[346] = "Advanced Algorithms"
	
		// Print the map
		fmt.Println(modules)
}
