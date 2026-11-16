
package main

import "fmt"

func main() {
	// Variables using var
	var studentName string = "Shrinivas"
	var age int = 23
	var temperature float64 = 30.5
	var isLearningGo bool = true

	// Short variable declaration
	city := "Chennai"

	// Constant
	const freezingPoint float64 = 0

	fmt.Println("===== Go Lang Day 1 =====")
	fmt.Println("Student:", studentName)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Temperature:", temperature)
	fmt.Println("Learning Go:", isLearningGo)
	fmt.Println("Freezing Point:", freezingPoint)

	// Celsius to Fahrenheit converter
	var celsius float64

	fmt.Print("\nEnter temperature in Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32

	fmt.Printf("%.2f Celsius = %.2f Fahrenheit\n",
		celsius, fahrenheit)
}
