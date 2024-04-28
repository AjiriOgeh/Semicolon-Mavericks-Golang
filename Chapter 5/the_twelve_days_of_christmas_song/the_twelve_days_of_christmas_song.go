package main

import "fmt"

func main() {
	var day int = 1

	for day <= 12 {
		switch day {
		case 1: fmt.Println("On the first day of Christmas")
		case 2: fmt.Println("On the second day of Christmas")
		case 3: fmt.Println("On the third day of Christmas")
		case 4: fmt.Println("On the fourth day of Christmas")
		case 5: fmt.Println("On the fifth day of Christmas")
		case 6: fmt.Println("On the sixth day of Christmas")
		case 7: fmt.Println("On the seventh day of Christmas")
		case 8: fmt.Println("On the eight day of Christmas")
		case 9: fmt.Println("On the ninth day of Christmas")
		case 10: fmt.Println("On the tenth day of Christmas")
		case 11: fmt.Println("On the eleventh day of Christmas")
		case 12: fmt.Println("On the twelfth day of Christmas")

		}

		fmt.Println("My true love sent to me:")

		switch day {
		case 1: fmt.Println("A Partridge in a Pear Tree.")
				fmt.Println()

		case 2: fmt.Println("Two Turtle Doves")
				fmt.Println("And a Partridge in a Pear Tree.")
				fmt.Println()

		case 3: fmt.Println("Three French Hens")
				fmt.Println("Two Turtle Doves,")
				fmt.Println("And a Partridge in a Pear Tree.")
				fmt.Println()

		case 4: fmt.Println("Four Calling Birds")
				fmt.Println("Three French Hens,")
				fmt.Println("Two Turtle Doves,")
				fmt.Println("And a Partridge in a Pear Tree.")
				fmt.Println()

		case 5: fmt.Println("Five Golden Rings")
				fmt.Println("Four Calling Birds,")
				fmt.Println("Three French Hens,")
				fmt.Println("Two Turtle Doves,")
				fmt.Println("And a Partridge in a Pear Tree.")
				fmt.Println()

		case 6:	fmt.Println("Six Geese a Laying,")
				fmt.Println("Five Golden Rings,")
				fmt.Println("Four Calling Birds,")
				fmt.Println("Three French Hens,")
				fmt.Println("Two Turtle Doves,")
				fmt.Println("And a Partridge in a Pear Tree.")
				fmt.Println()

		case 7: fmt.Println("Seven Swans a Swimming")
				fmt.Println("Six Geese a Laying,")
				fmt.Println("Five Golden Rings,")
				fmt.Println("Four Calling Birds,")
				fmt.Println("Three French Hens,")
				fmt.Println("Two Turtle Doves,")
				fmt.Println("And a Partridge in a Pear Tree.")
				fmt.Println()

		case 8: fmt.Println("Eight Maids a Milking")
				fmt.Println("Seven Swans a Swimming,")
				fmt.Println("Six Geese a Laying,")
				fmt.Println("Five Golden Rings,")
				fmt.Println("Four Calling Birds,")
				fmt.Println("Three French Hens,")
				fmt.Println("Two Turtle Doves,")
				fmt.Println("And a Partridge in a Pear Tree.")
				fmt.Println()

		case 9: fmt.Println("Nine Ladies Dancing")
				fmt.Println("Eight Maids a Milking,")
				fmt.Println("Seven Swans a Swimming,")
				fmt.Println("Six Geese a Laying,")
				fmt.Println("Five Golden Rings,")
				fmt.Println("Four Calling Birds,")
				fmt.Println("Three French Hens,")
				fmt.Println("Two Turtle Doves,")
				fmt.Println("And a Partridge in a Pear Tree.")
				fmt.Println()

		case 10: fmt.Println("Ten Lords a Leaping")
				 fmt.Println("Nine Ladies Dancing,")
				 fmt.Println("Eight Maids a Milking,")
				 fmt.Println("Seven Swans a Swimming,")
				 fmt.Println("Six Geese a Laying,")
				 fmt.Println("Five Golden Rings,")
				 fmt.Println("Four Calling Birds,")
				 fmt.Println("Three French Hens,")
				 fmt.Println("Two Turtle Doves,")
				 fmt.Println("And a Partridge in a Pear Tree.")
				 fmt.Println()

		case 11: fmt.Println("Eleven Pipers Piping")
				 fmt.Println("Ten Lords a Leaping,")
				 fmt.Println("Nine Ladies Dancing,")
				 fmt.Println("Eight Maids a Milking,")
				 fmt.Println("Seven Swans a Swimming,")
				 fmt.Println("Six Geese a Laying,")
				 fmt.Println("Five Golden Rings,")
				 fmt.Println("Four Calling Birds,")
				 fmt.Println("Three French Hens,")
				 fmt.Println("Two Turtle Doves,")
				 fmt.Println("And a Partridge in a Pear Tree.")
				 fmt.Println()

		case 12: fmt.Println("Twelve Drummers Drumming")
				 fmt.Println("Eleven Pipers Piping,")
				 fmt.Println("Ten Lords a Leaping,")
				 fmt.Println("Nine Ladies Dancing,")
				 fmt.Println("Eight Maids a Milking,")
				 fmt.Println("Seven Swans a Swimming,")
				 fmt.Println("Six Geese a Laying,")
				 fmt.Println("Five Golden Rings,")
				 fmt.Println("Four Calling Birds,")
				 fmt.Println("Three French Hens,")
				 fmt.Println("Two Turtle Doves,")
				 fmt.Println("And a Partridge in a Pear Tree.")
				 fmt.Println()
		}
		day++
	}
}