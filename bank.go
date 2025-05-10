package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("-----")
		// panic("failed to get balance")
	}

	fmt.Println("Welcome to Go bank!")

	fmt.Println("Dummy phone no", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int

		fmt.Print("Select one: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Deposit Successful. Your new balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				println("Invalid amount. You can't withdraw more than you have")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Withdrawal successful. Your new balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			// break
			return

		}

		// if choice == 1 {
		// 	fmt.Println("Your account balance is", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		// return
		// 		continue
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("Deposit Successful. Your new balance is:", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Withdraw amount: ")
		// 	var withdrawAmount float64

		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		return
		// 	}

		// 	if withdrawAmount > accountBalance {
		// 		println("Invalid amount. You can't withdraw more than you have")
		// 		return
		// 	}

		// 	fmt.Scan(&withdrawAmount)
		// 	accountBalance -= withdrawAmount
		// 	fmt.Println("Withdrawal successful. Your new balance is:", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// 	// return
		// }
	}
}
