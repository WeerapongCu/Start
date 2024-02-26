package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expactedReturnRate float64

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expactedReturnRate)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculatrFutureValue(investmentAmount, expactedReturnRate, years)
	//futureValue := investmentAmount * math.Pow(1 + expactedReturnRate / 100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Valur: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (adjusted for Inflation)", futureRealValue)
	// fmt.Printf(`Future Valur: %.1f\n

	// Future Value (adjusted for Inflation): %.1f`,futureValue,futureRealValue)
	//fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculatrFutureValue(inverstmentAmount, expactedReturnRate, years float64)(float64,float64){
	fv := inverstmentAmount * math.Pow(1 + expactedReturnRate / 100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}