package main

import (
	"fmt"
	"math"
)

func main() {

	var counter int
	var largest int = math.MinInt64 + 1
	var secondLargest int = math.MinInt64

	for counter < 10 {
		var number int
		fmt.Print("Enter a number: ")
		fmt.Scanf("%d\n", &number)

		if number > largest {
			temp := largest
			largest = number
			secondLargest = temp
		} else if number > secondLargest && number != largest{
			secondLargest = number
		}
		counter++
	}

	fmt.Printf("The largest number is %d\n", largest)
	fmt.Printf("The second largest number is %d\n", secondLargest)
}