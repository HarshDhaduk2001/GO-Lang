package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	storeResults(ebt, profit, ratio)

	fmt.Println("Earning before tax:", ebt)
	fmt.Println("Earning after tax:", profit)
	fmt.Println("Ratio:", ratio)
}

func getUserInput(infoText string) (float64, error) {
	var useInput float64
	fmt.Print(infoText)
	fmt.Scan(&useInput)

	if useInput <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return useInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (eBT float64, profit float64, ratio float64) {
	eBT = revenue - expenses
	profit = eBT * (1 - taxRate/100)
	ratio = eBT / profit
	return eBT, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
