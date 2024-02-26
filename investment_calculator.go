package main

import (
	"fmt"
	"math"
)

func main() {
	const intflationRate = 2.5
	var investmentAmount float64
	years := 10.0
	expactedReturnRate := 5.5

	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1 + expactedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1+intflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
