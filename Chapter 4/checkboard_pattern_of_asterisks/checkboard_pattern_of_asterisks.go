package main 

import "fmt"

func main() {
	
	var firstCounter int = 1

	for firstCounter <= 8 {
		if firstCounter % 2 == 0 {
			fmt.Print(" ")
		}
		var secondCounter = 1
		for secondCounter <= 8 {
			fmt.Print("* ")
			secondCounter++
		}
		fmt.Println()
		firstCounter++
	}

}