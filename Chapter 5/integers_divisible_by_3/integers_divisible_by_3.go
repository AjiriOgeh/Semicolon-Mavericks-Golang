package main

import "fmt"

func main() {

	var divisibleBy3Total int

	for number := 1; number <= 30; number++ {
		if number % 3 == 0 {
			divisibleBy3Total += number
		}
	}
	
	fmt.Printf(`The sum of integers divisible by 3 between 1 and 30 is %d`, divisibleBy3Total)
}