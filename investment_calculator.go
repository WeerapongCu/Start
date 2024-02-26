package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000.0
	years := 10.0
	expactedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1 + expactedReturnRate / 100, years)
	fmt.Print(futureValue)
}
