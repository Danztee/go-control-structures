package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error reading balance file:", err)
		return 0, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("failed to parse balance")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceTxt), 0o644)
}

func main() {
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("-----")
		// panic("failed to get balance")
	}

	fmt.Println("Welcome to Go bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
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
