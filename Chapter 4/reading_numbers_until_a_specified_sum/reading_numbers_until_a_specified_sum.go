package main 

import "fmt" 

func main() {

	var baseNumber int
	var value int
	var sumOfValues int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&baseNumber)

	fmt.Print("Enter a value: ")
	fmt.Scan(&value)
	sumOfValues += value

	for sumOfValues < baseNumber {
		fmt.Print("Enter a value: ")
		fmt.Scan(&value)
		sumOfValues += value
	}

	if sumOfValues > baseNumber {
		fmt.Printf(`The sum of values %d is greater than %d`, sumOfValues, baseNumber)
	}else{
		fmt.Printf(`The sum of values is equal to %d`, baseNumber)
	}
}