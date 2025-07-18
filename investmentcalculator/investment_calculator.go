package main

import (
	"fmt"
	"math"
)

const InflationRate float64 = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment amount: ")
	_, err := fmt.Scan(&investmentAmount)
	if err != nil {
		return
	}

	fmt.Print("Return rate: ")
	_, err = fmt.Scan(&expectedReturnRate)
	if err != nil {
		return
	}

	fmt.Print("Years: ")
	_, err = fmt.Scan(&years)
	if err != nil {
		return
	}

	expectedReturn, realReturn := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Expected return: %.4f\n", expectedReturn)
	fmt.Printf("Return adjusted for inflation: %.4f\n", realReturn)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	expectedReturn := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	realReturn := expectedReturn / math.Pow(1+InflationRate/100, years)

	return expectedReturn, realReturn
}

// Alternative version with declaration in return type:
//func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (expectedReturn, realReturn float64) {
//	expectedReturn = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
//	realReturn = expectedReturn / math.Pow(1+InflationRate/100, years)
//
//	return
//}
