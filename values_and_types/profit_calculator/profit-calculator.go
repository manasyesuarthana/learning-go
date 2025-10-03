// package main

import (
	"fmt"
	"os"
)

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	banner := `
	 ____             __ _ _      ____      _            _       _             
	|  _ \ _ __ ___  / _(_) |_   / ___|__ _| | ___ _   _| | __ _| |_ ___  _ __ 
	| |_) | '__/ _ \| |_| | __| | |   / _  | |/ __| | | | |/ _  | __/ _ \| '__|
	|  __/| | | (_) |  _| | |_  | |__| (_| | | (__| |_| | | (_| | || (_) | |   
	|_|   |_|  \___/|_| |_|\__|  \____\__,_|_|\___|\__,_|_|\__,_|\__\___/|_|   
	`

	fmt.Println(banner)
	fmt.Println("We would need a couple of data from you first!")

	fmt.Print("What was your revenue (in USD): ")
	fmt.Scan(&revenue)

	fmt.Print("How much were your expenses (in USD): ")
	fmt.Scan(&expenses)

	fmt.Print("Current Tax Rate: ")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses
	profit := EBT - (EBT * taxRate/100) //earnings after tax
	
	ratio := EBT/profit

	fmt.Println()
	fmt.Println("===== RESULTS =====")
	// fmt.Println("Eearnings Before Tax: ", EBT)
	// fmt.Println("Profit after Tax: ", profit)
	// fmt.Printf("Ratio between EBT and Profit: %f\n", ratio)

	//to store a formatted string into a variable, we can use Sprintf
	//Other alternatives: Sprint and Sprintln will basically store the string we are about to print into a variable

	formattedEBT := fmt.Sprintf("Earnings before Tax: %.2f\n", EBT)
	formattedProfit := fmt.Sprintf("Profit after Tax: %.2f\n", profit)
	formattedRatio := fmt.Sprintf("Ratio Between EBT and Profit: %.2f\n", ratio)

	fmt.Print(formattedEBT, formattedProfit, formattedRatio)

	os.Exit(0)
	
}