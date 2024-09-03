package main

import "fmt"

func main() {

	useTaxDeduction := getTaxDeductionAmount()
	amountList := []float64{105.23, 2232.1357, 234.43, 2342.532, 56634.522}

	for idx, amount := range amountList {
		fmt.Printf("%v, %.3f\n", idx+1, useTaxDeduction(amount))
	}

}

func getTaxDeductionAmount() func(float64) float64 {

	tax := 18.0

	getDeductedTax := func(amount float64) float64 {
		var result float64 = (amount - (amount * (tax / 100)))
		return result
	}

	return getDeductedTax
}
