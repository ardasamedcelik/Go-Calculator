package main

import (
	"fmt"
)

func main() {
	// Calculator
	fmt.Println("___Calculator___")
	//Number 1
	fmt.Print("Enter the first number: ")
	var num1 int
	fmt.Scanln(&num1)

	//Number 2
	fmt.Print("Enter the second number: ")
	var num2 int
	fmt.Scanln(&num2)

	// User input
	var choice int

	fmt.Println("----------------")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("----------------")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	// Switch case
	switch choice {
	case 1:
		fmt.Println("Addition of two numbers is: ", num1+num2)

	case 2:
		fmt.Println("Subtraction of two numbers is: ", num1-num2)
	case 3:
		fmt.Println("Multiplication of two numbers is: ", num1*num2)
	case 4:
		fmt.Println("Division of two numbers is: ", num1/num2)
	case 5:
		fmt.Println("Exit")
	default:
		fmt.Println("Invalid choice")

	}
}
