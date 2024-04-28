package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		minimum = math.MaxInt64
		maximum = math.MinInt64

	)
	var numberOfValues int
	var number int

	fmt.Print("Enter the number of values you want to input: ")
	fmt.Scan(&numberOfValues)

	for count := 0; count < numberOfValues; count++ {
		
		fmt.Print("Enter a number: ")
		fmt.Scan(&number)

		if number > maximum {
			maximum = number
		}

		if number < minimum{
			minimum = number
		}

	}
	sumOfExtremes := minimum + maximum

	fmt.Print(`The sum of the two extremes is `, sumOfExtremes)
}