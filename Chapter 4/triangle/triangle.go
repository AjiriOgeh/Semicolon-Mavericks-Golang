package main

import "fmt"

func main(){

	var length int
	
	fmt.Print("Enter the triangle length: ")
	fmt.Scanln(&length)

	for length < 1 || length > 10 {
		fmt.Print("Enter the triangle length: ")
		fmt.Scanln(&length)
	}

	firstCounter := 1
	for firstCounter <= length {
		secondCounter := 1
		for secondCounter <= firstCounter {
			fmt.Print("*")
			secondCounter++
		}
		fmt.Println()
		firstCounter++
	}
}