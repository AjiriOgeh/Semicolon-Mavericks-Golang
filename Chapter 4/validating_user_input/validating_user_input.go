package main

import "fmt"

func main() {

	var (
		numberOfPasses = 0
		numberOfFailures = 0		
		count = 0
	)

	var result int

	for count < 10 {
		fmt.Print("Enter result (1 = pass, 2 = fail): ")
		fmt.Scanln(&result)
		
		for result < 1 || result > 2 {
			fmt.Print("Enter result (1 = pass, 2 = fail): ")
			fmt.Scanln(&result)
		}

		if result == 1 {
			numberOfPasses++
		} else {
			numberOfFailures++
		}
		count++
	}

	fmt.Println("Passed: ", numberOfPasses)
	fmt.Println("Failed: ", numberOfFailures)

	if (numberOfPasses > 8) {
		fmt.Println("Bonus to instructor!")
	}
}