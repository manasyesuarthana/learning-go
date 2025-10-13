package main

import "fmt"

func printMessage() {
	message := `
	Welcome to the Go Bank!
	What do you want to do?
	1. Check Balance
	2. Deposit Money
	3. Withdraw Money
	4. Exit
	`

	fmt.Println(message)
}
