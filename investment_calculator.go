package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 6.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("future value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("future value (adjusted for Inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	// return fv, rfv if return type is ( float64, float64) or only return if return specify value and type (fv float64, rfv float64)
	//return fv, rfv
	return
}
