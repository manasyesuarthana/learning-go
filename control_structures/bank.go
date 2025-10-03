package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func writeToFile(balance float64) {
	balanceText := fmt.Sprint("Current balance: ", balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFile)
	if err != nil {
		return 0, errors.New("failed to find balance.txt file")
	}
	parsedBalance, err := strconv.ParseFloat(string(data)[17:], 64)
	if err != nil {
		return 0, errors.New("failed to parse balance from file")
	}
	return parsedBalance, nil

}
func main() {
	message := `
	Welcome to the Go Bank!
	What do you want to do?
	1. Check Balance
	2. Deposit Money
	3. Withdraw Money
	4. Exit
	`
	var userChoice int
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("Initial Error: ", err)
		fmt.Println("Don't worry, depositing automatically creates a new one but starts from 0.")
	}

	isLoggedOn := true

	for isLoggedOn {
		fmt.Println(message)
		fmt.Print("Your choice: ")
		fmt.Scan(&userChoice)
		fmt.Println()

		switch userChoice {
		case 1:
			fmt.Printf("=> Your current balance is %.2f\n", accountBalance)

		case 2:
			var userDeposit float64
			fmt.Print("=> How much do you want to deposit: ")
			fmt.Scan(&userDeposit)
			accountBalance += userDeposit
			writeToFile(accountBalance)
			fmt.Printf("=> Thank you for your deposit. Your account balance is now: %.2f\n", accountBalance)

		case 3:
			var userWithdrawal float64
			isWithdrawed := false
			for !isWithdrawed {
				fmt.Print("=> Amount of money to withdraw from account: ")
				fmt.Scan(&userWithdrawal)
				if userWithdrawal > accountBalance {
					fmt.Printf("(!) Cannot bro. Not enough balance to withdraw.\n")
					fmt.Printf("Your current balance is %.2f\ne", accountBalance)
				} else if userWithdrawal <= 0 {
					fmt.Println("(!) You cannot withdraw nothing or below nothing...")
					fmt.Println("Please enter a valid amount.")
				} else {
					fmt.Println("=> Withdrawing...")
					fmt.Printf("=> You now have %.2f cash in your hands\n", userWithdrawal)
					accountBalance -= userWithdrawal
					writeToFile(accountBalance)
					fmt.Printf("=> Your balance is now %.2f\n", accountBalance)
					isWithdrawed = true
				}
			}

		case 4:
			fmt.Println("=> Thank you for your visit. Have a great day!")
			isLoggedOn = false

		default:
			fmt.Println("=> Please pick a valid option and NO HACKING >:( ")
		}
	}
	os.Exit(0)
}
