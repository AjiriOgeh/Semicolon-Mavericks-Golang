package checkout_system

import (
	"fmt"
	"time"
)

func FillTransactionInformation(checkoutPurchase *CheckoutPurchase) {
	var itemBought string
	var numberOfPieces int
	var pricePerUnit float64
	
	fmt.Print("What did the user buy?: ")
	fmt.Scanln(&itemBought)

	fmt.Print("How many pieces?: ")
	fmt.Scanln(&numberOfPieces)
	
	fmt.Print("How much per unit?: ")
	fmt.Scanln(&pricePerUnit)

	checkoutPurchase.ArrangePurchaseInformationToLists(itemBought, numberOfPieces, pricePerUnit)
	fmt.Println()
}

type CheckoutPurchase struct {
	CustomerName string	
	ItemsBoughtList []string
	NumberOfPiecesList []int
	PricePerUnitList []float64
	CashierName string
	Discount float64
	CustomerPayment float64
	Bill Bill
}

type Bill struct{
	SubTotal float64
	DiscountPrice float64
	BillTotal float64
	ValueAddedTax float64
	Balance float64
}

func (c *CheckoutPurchase) ArrangePurchaseInformationToLists(item string, numberOfPieces int, pricePerUnit float64) {
	c.ItemsBoughtList = append(c.ItemsBoughtList, item)
	c.NumberOfPiecesList = append(c.NumberOfPiecesList, numberOfPieces)
	c.PricePerUnitList = append(c.PricePerUnitList, pricePerUnit)
}

func (c *CheckoutPurchase) SetSubTotal() {
		for count := 0; count < len(c.ItemsBoughtList); count++{
		c.Bill.SubTotal += float64(c.NumberOfPiecesList[count]) * c.PricePerUnitList[count]
	}
}

func (c *CheckoutPurchase) SetDiscountPrice() {
	c.Bill.DiscountPrice = c.Bill.SubTotal * (c.Discount / 100)
}

func (c *CheckoutPurchase) SetValueAddedTax() {
	c.Bill.ValueAddedTax = (17.5 / 100) * c.Bill.SubTotal
}

func (c * CheckoutPurchase) SetBillTotal() {
	c.Bill.BillTotal = c.Bill.SubTotal + c.Bill.ValueAddedTax - c.Bill.DiscountPrice
}

func (c* CheckoutPurchase) SetBalance() {
	c.Bill.Balance = c.CustomerPayment - c.Bill.BillTotal
}

func DisplayHeader(checkoutPurchase CheckoutPurchase) string {
	transactionTime := time.Now()

	return fmt.Sprintf(`SEMICOLON STORES
MAIN BRANCH
LOCATION: 312, HERBERT MACAULAY WAY, SABO YABA, LAGOS.
TEL: 03293828343
Date: %v
Cashier: %s
Customer Name: %s
=========================================================
%5s %10s %10s %15s
---------------------------------------------------------
`,transactionTime,checkoutPurchase.CashierName, checkoutPurchase.CustomerName, "ITEM", "QTY", "PRICE", "TOTAL(NGN)")
}

func DisplayItemsBoughtWithPrices(checkoutPurchase CheckoutPurchase) {
	for count := 0; count < len(checkoutPurchase.ItemsBoughtList); count++ {
		fmt.Printf("%5s %10d %10.2f %15.2f\n",checkoutPurchase.ItemsBoughtList[count], 
		checkoutPurchase.NumberOfPiecesList[count], checkoutPurchase.PricePerUnitList[count], 
		(float64(checkoutPurchase.NumberOfPiecesList[count]) * checkoutPurchase.PricePerUnitList[count]))
	}
	fmt.Println("---------------------------------------------------------")
}

func SetBillInformation(checkoutPurchase *CheckoutPurchase) {
	checkoutPurchase.SetSubTotal()
	checkoutPurchase.SetDiscountPrice()
	checkoutPurchase.SetValueAddedTax()
	checkoutPurchase.SetBillTotal()
}

func SetBillTotalInformationAfterPayment(checkoutPurchase *CheckoutPurchase) {
	checkoutPurchase.SetBalance()
}

func DisplayExpenses(checkoutPurchase CheckoutPurchase) string{
	return fmt.Sprintf(`%30s : %8.2f
%30s : %8.2f
%30s : %8.2f
=========================================================`, "Sub Total", checkoutPurchase.Bill.SubTotal, "Discount", 
checkoutPurchase.Bill.DiscountPrice, "VAT @ 17.50%", checkoutPurchase.Bill.ValueAddedTax)
}

func DisplayBillTotal(checkoutPurchase CheckoutPurchase) string{
	return fmt.Sprintf(`%30s : %8.2f
=========================================================
THIS IS NOT A RECEIPT. KINDLY PAY %.2f
=========================================================
`, "Bill Total", checkoutPurchase.Bill.BillTotal, checkoutPurchase.Bill.BillTotal)
}

func DisplayBalanceAfterPayment(checkoutPurchase CheckoutPurchase) string{
	return fmt.Sprintf(`%30s : %8.2f
%30s : %8.2f
%30s : %8.2f
=========================================================
THANK YOU FOR YOU PATRONAGE
=========================================================`, "Bill Total", checkoutPurchase.Bill.BillTotal, "Amount Paid", checkoutPurchase.CustomerPayment,
 "Balance", checkoutPurchase.Bill.Balance)
}



