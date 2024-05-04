package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		// panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount.")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amont. You can't withdraw more that your account balance.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
		default:
			fmt.Println("Good bye!")
			fmt.Println("Thanks for choosing our bank.")
			return
		}
	}
}
