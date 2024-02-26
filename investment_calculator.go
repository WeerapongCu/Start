package main

import (
	"fmt"
	"math"
)

func main() {
	const intflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expactedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expactedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1 + expactedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1+intflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Valur: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (adjusted for Inflation)", futureRealValue)
	//fmt.Printf("Future Valur: %.0f\nFuture Value (adjusted for Inflation): %.0f",futureValue,futureRealValue)
	//fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}
