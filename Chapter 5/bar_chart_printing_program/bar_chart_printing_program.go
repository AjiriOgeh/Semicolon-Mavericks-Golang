package main 

import "fmt"

func main() {
	var firstNumber int
	var secondNumber int
	var thirdNumber int
	var fourthNumber int
	var fifthNumber int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&firstNumber)


	for firstNumber < 1 || firstNumber > 30 {
		fmt.Print("Enter the first number: ")
		fmt.Scanln(&firstNumber)
	}

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&secondNumber)

	for secondNumber < 1 || secondNumber > 30 {
		fmt.Print("Enter the second number: ")
		fmt.Scanln(&secondNumber)
	}

	fmt.Print("Enter the third number: ")
	fmt.Scanln(&thirdNumber)

	for thirdNumber < 1 || thirdNumber > 30 {
		fmt.Print("Enter the third number: ")
		fmt.Scanln(&thirdNumber)
	}

	fmt.Print("Enter the fourth number: ")
	fmt.Scanln(&fourthNumber)

	for fourthNumber < 1 || fourthNumber > 30 {
		fmt.Print("Enter the fourth number: ")
		fmt.Scanln(&fourthNumber)
	}

	fmt.Print("Enter the five number: ")
	fmt.Scanln(&fifthNumber)

	for fifthNumber < 1 || fifthNumber > 30 {
		fmt.Print("Enter the five number: ")
		fmt.Scanln(&fifthNumber)
	}

	for count := 0; count < firstNumber; count++{
		fmt.Print(`*`)
	}
	fmt.Println()

	for count := 0; count < secondNumber; count++{
		fmt.Print(`*`)
	}
	fmt.Println()

	for count := 0; count < thirdNumber; count++{
		fmt.Print(`*`)
	}
	fmt.Println()

	for count := 0; count < fourthNumber; count++{
		fmt.Print(`*`)
	}
	fmt.Println()

	for count := 0; count < fifthNumber; count++{
		fmt.Print(`*`)
	}
}