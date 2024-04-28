package main 

import (
	"fmt"
	"strconv"
)

func main () {
	var data string 
	var decryptedData string

	fmt.Print("Enter the data to decrypt: ")
	fmt.Scanf("%s\n", &data)

	for len(data) != 4 {
		fmt.Print("Enter the data to decrypt: ")
		fmt.Scanf("%s\n", &data)
	}

	dataFirstNumber, _ := strconv.Atoi(string(data[0]))
	dataSecondNumber, _ := strconv.Atoi(string(data[1]))
	dataThirdNumber, _ := strconv.Atoi(string(data[2]))
	dataFourthNumber, _ := strconv.Atoi(string(data[3]))

	decryptedDataFirstNumber := (dataThirdNumber + 3) % 10
	decryptedDataSecondNumber := (dataFourthNumber + 3) % 10
	decryptedDataThirdNumber := (dataFirstNumber + 3) % 10
	decryptedDataFourthNumber := (dataSecondNumber +3) % 10

	decryptedData += strconv.Itoa(decryptedDataFirstNumber) + strconv.Itoa(decryptedDataSecondNumber) + strconv.Itoa(decryptedDataThirdNumber) + strconv.Itoa(decryptedDataFourthNumber)

	fmt.Printf("Encrypted data: '%s' --> Decrypted:[%s]", data, decryptedData)
}

