package main

import "fmt"

func main(){

	number := 1
	fmt.Printf("N\tN2\tN3\tN4\n")

	for number <= 5 {
		square := number * number
		cube := number * number * number
		fourthPower := number * number * number * number
		fmt.Printf("%d\t%d\t%d\t%d\n", number, square, cube, fourthPower,)
		number++
	}
}