package main

import "fmt"

// func main() {
// 	// Array
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
// 	fmt.Println(prices)

// 	var productNames [4]string
// 	productNames = [4]string{"Apple", "Bat"}
// 	productNames[2] = "Cat"
// 	fmt.Println(productNames[1])
// 	fmt.Println(productNames)
// 	fmt.Println(len(productNames), cap(productNames)) // len->length, cap->capacity

// 	//Slice
// 	featuredPrices := prices[1:3]
// 	fmt.Println(featuredPrices)

// 	featuredPrices1 := prices[:3]
// 	fmt.Println(featuredPrices1)

// 	featuredPrices2 := prices[1:]
// 	fmt.Println(featuredPrices2)

// }

func main() {
	prices := []float64{10.99, 8.99}
	prices[1] = 9.99
	fmt.Println(prices)

	prices = append(prices, 5.99)
	fmt.Println(prices)

	discountPrice := []float64{101.99, 122.00, 13.99}
	prices = append(prices, discountPrice...)
	fmt.Println(prices)
}
