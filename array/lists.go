package main

import (
	"fmt"
)

func main() {
	var productNames [5]string = [5]string{"A book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	fmt.Println(prices[2])
	fmt.Println(productNames)

	featuredPrices := prices[:3]
	fmt.Println(featuredPrices)
}