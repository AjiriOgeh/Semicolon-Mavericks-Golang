package main

import "fmt"

func main() {
	
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	for number < 10000 || number > 99999 {
		fmt.Print("Enter a number: ")
		fmt.Scanln(&number)
	}

	reversedNumber := 0
	modifiedNumber := number 
	count := 0

	for count < 5 {
		reversedNumber = (reversedNumber* 10) + (modifiedNumber % 10)
		modifiedNumber /= 10		
		count++
	}

	if number == reversedNumber{
		fmt.Printf("%d is a palindrome number\n", number)
	}else {
		fmt.Printf("%d is not a palindrome number\n", number)
	}
}