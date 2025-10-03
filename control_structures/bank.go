package main

import (
	"fmt"
	"os"
)

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
	var accountBalance float64 = 1000 //dummy balance
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
