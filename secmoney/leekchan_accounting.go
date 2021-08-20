package main

import (
	"fmt"
	"math/big"

	"github.com/leekchan/accounting"
)

// testleekchanaccounting test common issues with float with go using accouting package
// this package inside contains shopspring/decimal tested bellow
func testleekchanaccounting() {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213123))                       // "$123,456,789.21"
	fmt.Println(ac.FormatMoney(12345678))                               // "$12,345,678.00"
	fmt.Println(ac.FormatMoney(big.NewRat(77777777, 3)))                // "$25,925,925.67"
	fmt.Println(ac.FormatMoney(big.NewRat(-77777777, 3)))               // "-$25,925,925.67"
	fmt.Println(ac.FormatMoneyBigFloat(big.NewFloat(123456789.213123))) // "$123,456,789.21"
	//fmt.Println(ac.FormatMoneyDecimal(decimal.New(123456789.213123, 0))) // "$123,456,789.21"

	ac = accounting.Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
	fmt.Println(ac.FormatMoney(4999.99))  // "€4.999,99"

	fmt.Println(ac.FormatMoney(500000)) // "$500,000.00"

	ac = accounting.Accounting{Symbol: "£ ", Precision: 0}
	fmt.Println(ac.FormatMoney(500000)) // "£ 500,000"

	ac = accounting.Accounting{Symbol: "GBP", Precision: 0,
		Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
	fmt.Println(ac.FormatMoney(1000000)) // "GBP 1,000,000"
	fmt.Println(ac.FormatMoney(-5000))   // "GBP (5,000)"
	fmt.Println(ac.FormatMoney(0))       // "GBP --"
}