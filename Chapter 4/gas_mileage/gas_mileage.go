package main

import "fmt"

func main() {

	var gallonsUsed float32
	var milesDriven float32
	var totalGallonsUsed float32
	var totalMilesDriven float32

	fmt.Print("Enter the gallons used (-1 to end): ")
	fmt.Scanf("%f\n", &gallonsUsed)

	fmt.Print("Enter the miles driven: ")
	fmt.Scanf("%f\n", &milesDriven)

	for gallonsUsed != -1 {
		fmt.Println("The miles/ gallon for this tank was ", (milesDriven / gallonsUsed))
		fmt.Println()

		totalGallonsUsed += gallonsUsed
		totalMilesDriven += milesDriven

		fmt.Print("Enter the gallons used (-1 to end): ")
		fmt.Scanf("%f\n", &gallonsUsed)
	
		if gallonsUsed == -1 {
			break
		}
	
		fmt.Print("Enter the miles driven: ")
		fmt.Scanf("%f\n", &milesDriven)
	}

	combinedMilesPerGallon := totalMilesDriven / totalGallonsUsed

	fmt.Println()
	fmt.Printf("The overall average miles/gallon was %f", combinedMilesPerGallon)
}