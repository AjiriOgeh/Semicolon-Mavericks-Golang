package main

import (
	"fmt"
	"strconv"
)

func main() {
	var data string 
	var encryptedData string

	fmt.Print("Enter the data to encrypt: ")
	fmt.Scanf("%s\n", &data)

	for len(data) != 4 {
		fmt.Print("Enter the data to encrypt: ")
		fmt.Scanf("%s\n", &data)
	}

	dataFirstNumber, _ := strconv.Atoi(string(data[0]))
	dataSecondNumber, _ := strconv.Atoi(string(data[1]))
	dataThirdNumber, _ := strconv.Atoi(string(data[2]))
	dataFourthNumber, _ := strconv.Atoi(string(data[3]))

	encryptedDataFirstNumber := (dataThirdNumber + 7) % 10
	encryptedDataSecondNumber := (dataFourthNumber + 7) % 10
	encryptedDataThirdNumber := (dataFirstNumber + 7) % 10
	encryptedDataFourthNumber := (dataSecondNumber + 7) % 10

	encryptedData += strconv.Itoa(encryptedDataFirstNumber) + strconv.Itoa(encryptedDataSecondNumber) + strconv.Itoa(encryptedDataThirdNumber) + strconv.Itoa(encryptedDataFourthNumber)

	fmt.Printf("Data: '%s' --> Encrypted:[%s]", data, encryptedData)
}