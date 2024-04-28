package main

import "fmt"

func main() {
	fmt.Println("(A)")
		
	for count := 1; count <= 10; count++{
		for counter := 1; counter <= count; counter++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("(B)")
	for count := 10; count >= 1; count-- {
		for counter := 1; counter <= count; counter++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("(C)")
	numberOfBlankSpaces := 0
	for count := 10; count >= 1; count-- {
		for i := 0; i < numberOfBlankSpaces; i++ {
			fmt.Print(" ")
		}

		for counter := 1; counter <= count; counter++ {
			fmt.Print("*")
		}
		fmt.Println()
		numberOfBlankSpaces++
	}
	fmt.Println()

	numberOfBlankSpaces = 0
	fmt.Println("(D)")
	for count := 1; count <= 10; count++{
		for i := 9; i > numberOfBlankSpaces; i-- {
			fmt.Print(" ")
		}

		for counter := 1; counter <= count; counter++{
			fmt.Print("*")
		}
		numberOfBlankSpaces++
		fmt.Println()
	}
}