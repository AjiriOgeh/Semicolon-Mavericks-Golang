package main

import "fmt"

func main() {

	var firstNumber int
	var secondNumber int

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&firstNumber)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&secondNumber)

	if firstNumber > secondNumber {
		fmt.Print(1)
	}else if firstNumber < secondNumber{
		fmt.Print(-1)
	}else {
		fmt.Print(0)
	}
}