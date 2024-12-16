package main

import (
	"fmt"
)

func convertCelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	var choice int
	fmt.Println("Choose an option:")
	fmt.Println("1. Convert Celsius to Fahrenheit")
	fmt.Println("2. Convert Fahrenheit to Celsius")
	fmt.Print("Enter your choice (1 or 2): ")
	fmt.Scan(&choice)

	if choice == 1 {
		var celsius float64
		fmt.Print("Enter temperature in Celsius: ")
		fmt.Scan(&celsius)
		fahrenheit := convertCelsiusToFahrenheit(celsius)
		fmt.Printf("%.2f°C is equal to %.2f°F\n", celsius, fahrenheit)
	} else if choice == 2 {
		var fahrenheit float64
		fmt.Print("Enter temperature in Fahrenheit: ")
		fmt.Scan(&fahrenheit)
		celsius := convertFahrenheitToCelsius(fahrenheit)
		fmt.Printf("%.2f°F is equal to %.2f°C\n", fahrenheit, celsius)
	} else {
		fmt.Println("Invalid choice. Please enter 1 or 2.")
	}
}
