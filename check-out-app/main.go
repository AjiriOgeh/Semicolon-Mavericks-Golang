package main

import (
	"fmt"
	"check-out-app/checkout_system"
)

func main() {
	var checkoutPurchase checkout_system.CheckoutPurchase
	var moreItems string

	fmt.Print("What is the customer's name: ")
	fmt.Scanln(&checkoutPurchase.CustomerName)

	checkout_system.FillTransactionInformation(&checkoutPurchase)

	fmt.Print("Add more items?: ")
	fmt.Scanln(&moreItems)
	fmt.Println()

	for moreItems == "yes" {
		checkout_system.FillTransactionInformation(&checkoutPurchase)

		fmt.Print("Add more items?: ")
		fmt.Scanln(&moreItems)
		fmt.Println()
	}

	fmt.Print("What is your name? ")
	fmt.Scanln(&checkoutPurchase.CashierName)

	fmt.Print("How much discount will the coustomer get? ")
	fmt.Scanln(&checkoutPurchase.Discount)
	fmt.Println()
	fmt.Println()

	fmt.Println(checkout_system.DisplayHeader(checkoutPurchase))	
	checkout_system.DisplayItemsBoughtWithPrices(checkoutPurchase)

	checkout_system.SetBillInformation(&checkoutPurchase)

	fmt.Println(checkout_system.DisplayExpenses(checkoutPurchase))
	fmt.Println(checkout_system.DisplayBillTotal(checkoutPurchase))
	fmt.Println()
	fmt.Println()

	fmt.Print("How much did the customer give you? ")
	fmt.Scanln(&checkoutPurchase.CustomerPayment)
	fmt.Println()
	fmt.Println()

	checkout_system.SetBillTotalInformationAfterPayment(&checkoutPurchase)

	fmt.Println(checkout_system.DisplayHeader(checkoutPurchase))	
	checkout_system.DisplayItemsBoughtWithPrices(checkoutPurchase)
	fmt.Println(checkout_system.DisplayExpenses(checkoutPurchase))
	fmt.Println(checkout_system.DisplayBillTotal(checkoutPurchase))
	fmt.Println(checkout_system.DisplayBalanceAfterPayment(checkoutPurchase))
}