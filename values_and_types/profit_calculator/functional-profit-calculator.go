package main

import (
	"errors"
	"fmt"
	"os"
)

// Task: add error handling
// 1) Validate User Input
//		=> Show error message & exit if invalid input is provided: negative, non-integer, 0
// 2) Store calculated values into file.

var userRevenue float64
var userExpenses float64
var userTaxRate float64

const resultFile = "results.txt"

func main() {
	printBanner()
	askUserRevenue()
	askUserExpense()
	askUserTaxRate()

	userEBT, userProfit, userRatio := calculateProfit(userRevenue, userExpenses, userTaxRate)
	// printResults(userEBT, userProfit, userRatio)
	writeResultToFile(userEBT, userProfit, userRatio)

	os.Exit(0)
}

func askUserRevenue() {
	fmt.Print("What was your revenue: ")
	fmt.Scan(&userRevenue)

	err := validateInput(userRevenue)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func askUserExpense() {
	fmt.Print("What was the total of your expenses: ")
	fmt.Scan(&userExpenses)

	err := validateInput(userExpenses)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func askUserTaxRate() {
	fmt.Print("What is your current tax rate percentage: ")
	fmt.Scan(&userTaxRate)

	err := validateInput(userTaxRate)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt - (ebt * taxRate / 100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

// func printResults(ebt, profit, ratio float64) {
// 	fmt.Printf("\n")
// 	fmt.Println("==== Your Results ====")
// 	fmt.Printf("Earnings before Tax: %.2f\n", ebt)
// 	fmt.Printf("Profit after Tax: %.2f\n", profit)
// 	fmt.Printf("Ratio Between EBT and Profit: %.2f\n", ratio)
// 	fmt.Printf("\n")
// }

func writeResultToFile(EBT, profit, ratio float64) {
	result := fmt.Sprintf(`
	==== Your Results ====
	Earnings before Tax: %.2f
	Profit after Tax: %.2f
	Ratio between EBT and Profit: %.2f
	`, EBT, profit, ratio)

	os.WriteFile(resultFile, []byte(result), 0644)
	fmt.Printf("Calculation Complete. Results written to %s.\nPlease check the contents of the file.\n", resultFile)

}

func validateInput(input float64) error {

	// inputType := reflect.TypeOf(input).String() - this is uncessary since go automatically converts non numbers to 0

	if input <= 0 {
		return errors.New("Please enter a valid input: no zeros, negatives, or strings.")
	}
	return nil
}

func printBanner() {
	banner := `
	 ____             __ _ _      ____      _            _       _             
	|  _ \ _ __ ___  / _(_) |_   / ___|__ _| | ___ _   _| | __ _| |_ ___  _ __ 
	| |_) | '__/ _ \| |_| | __| | |   / _  | |/ __| | | | |/ _  | __/ _ \| '__|
	|  __/| | | (_) |  _| | |_  | |__| (_| | | (__| |_| | | (_| | || (_) | |   
	|_|   |_|  \___/|_| |_|\__|  \____\__,_|_|\___|\__,_|_|\__,_|\__\___/|_|   
	`
	fmt.Println(banner)
}
