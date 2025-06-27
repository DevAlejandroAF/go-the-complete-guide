package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Println("======== Investment Calculator ========")
	fmt.Print("Investment amount: ")
	investmentInput, _ := reader.ReadString('\n')
	investmentAmount, _ = strconv.ParseFloat(strings.TrimSpace(investmentInput), 64)

	fmt.Print("Expected return rate: ")
	returnRateInput, _ := reader.ReadString('\n')
	expectedReturnRate, _ = strconv.ParseFloat(strings.TrimSpace(returnRateInput), 64)

	fmt.Print("Years: ")
	yearsInput, _ := reader.ReadString('\n')
	years, _ = strconv.ParseFloat(strings.TrimSpace(yearsInput), 64)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future value: %f\n", futureValue)
	fmt.Printf("Future real value: %f", futureRealValue)
}
