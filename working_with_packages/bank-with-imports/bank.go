package main

import (
	"fmt"
	"os"

	"example.com/bank-with-imports/fileio"
)

func main() {

	var userChoice int
	accountBalance, err := fileio.GetBalanceFromFile()

	if err != nil {
		fmt.Println("Initial Error: ", err)
		fmt.Println("Don't worry, depositing automatically creates a new one but starts from 0.")
	}

	isLoggedOn := true

	for isLoggedOn {
		printMessage() //since message.go also has package main, we dont need to import it here anymore. We can just call it directly. But we have to run it with `go run .`
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
			fileio.WriteToFile(accountBalance)
			fmt.Printf("=> Thank you for your deposit. Your account balance is now: %.2f\n", accountBalance)

		case 3:
			var userWithdrawal float64

			isWithdrawed := false //keep asking the user for withdraw amount until they pass in a valid input.
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
					fileio.WriteToFile(accountBalance)
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
