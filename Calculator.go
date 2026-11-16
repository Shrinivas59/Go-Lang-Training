
package main

import "fmt"

// Function with multiple return values
func add(a, b float64) (float64, string) {
	return a + b, "Addition successful"
}

// Function with named return values
func divide(a, b float64) (result float64, message string) {
	if b == 0 {
		message = "Error: Division by zero is not allowed"
		return
	}

	result = a / b
	message = "Division successful"
	return
}

func main() {
	defer fmt.Println("\nCalculator program completed!")

	fmt.Println("===== Go Lang Day 2 =====")

	// If / Else
	number := 10

	if number > 0 {
		fmt.Println("Number is positive")
	} else if number < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is zero")
	}

	// Switch statement
	day := 2

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Other day")
	}

	// Standard for loop
	fmt.Println("\nFor loop:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// While-style loop
	fmt.Println("\nWhile-style loop:")
	count := 1

	for count <= 3 {
		fmt.Println(count)
		count++
	}

	// Calculator
	var a, b float64
	var operator string

	fmt.Print("\nEnter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	var result float64
	var message string

	switch operator {
	case "+":
		result, message = add(a, b)
	case "-":
		result = a - b
		message = "Subtraction successful"
	case "*":
		result = a * b
		message = "Multiplication successful"
	case "/":
		result, message = divide(a, b)
	default:
		message = "Invalid operator"
	}

	fmt.Println(message)

	if operator == "+" || operator == "-" ||
		operator == "*" || operator == "/" {
		if operator == "/" && b == 0 {
			return
		}
		fmt.Println("Result:", result)
	}
}
