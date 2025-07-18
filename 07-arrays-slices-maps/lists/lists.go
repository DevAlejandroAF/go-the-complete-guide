package lists

import "fmt"

func list() {
	var productNames [4]string = [4]string{"A book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A carpet"
	fmt.Println(productNames)
	fmt.Println(prices[2])

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
