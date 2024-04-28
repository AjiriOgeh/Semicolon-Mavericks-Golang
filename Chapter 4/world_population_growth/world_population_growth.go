package main

import "fmt"

func main() {
	var currentPopulation int = 8_106_303_120
	const worldPopulationGrowth float32 = 0.0091

	fmt.Printf("%s\t%s\t%s\n", "Year", "NewPopulation", "Difference")

	for year := 1; year <= 75; year++ {
		newPopulation := int(float32(currentPopulation) * (1 + worldPopulationGrowth))
		difference := newPopulation - currentPopulation


		fmt.Printf("%2d\t%8d\t%5d\n", year, newPopulation, difference)
		currentPopulation = int(newPopulation)
	}
}
