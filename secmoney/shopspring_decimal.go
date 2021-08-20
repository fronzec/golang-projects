package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func testDecimal() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875

	var n= decimal.NewFromFloat(0.0) // Number
	var accumulator= decimal.NewFromFloat(0.0) // Number
	for i := 0; i < 1000; i++ {
		increment := decimal.NewFromFloat(.01)
		n.Add(increment) // n is not affected by operation, we need to reasign as above
		accumulator = accumulator.Add(increment)
	}
	fmt.Println("Increment:", n) // Increment is not affected by operations
	fmt.Println("Acummulator:", accumulator) // Acummulator is successfully incremented
	a:= decimal.NewFromFloat(.333)
	b:= decimal.NewFromFloat(.333)
	c:= decimal.NewFromFloat(.333)
	r := a.Add(b).Add(c)
	fmt.Println("C:", r) // Acummulator is successfully incremented
}