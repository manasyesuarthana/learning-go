package main

import (
	"fmt"
	"os"
)

var userRevenue float64
var userExpenses float64
var userTaxRate float64

func main() {
	printBanner()
	askUserRevenue()
	askUserExpense()
	askUserTaxRate()

	userEBT, userProfit, userRatio := calculateProfit(userRevenue, userExpenses, userTaxRate)
	printResults(userEBT, userProfit, userRatio)

	os.Exit(0)
}

func askUserRevenue() {
	fmt.Print("What was your revenue: ")
	fmt.Scan(&userRevenue)
}

func askUserExpense() {
	fmt.Print("What was the total of your expenses: ")
	fmt.Scan(&userExpenses)
}

func askUserTaxRate() {
	fmt.Print("What is your current tax rate percentage: ")
	fmt.Scan(&userTaxRate)
}

func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64){
	ebt := 	revenue - expenses
	profit := ebt - (ebt * taxRate/100)
	ratio := ebt/profit
	return ebt, profit, ratio
}

func printResults(ebt, profit, ratio float64) {
	fmt.Printf("\n")
	fmt.Println("==== Your Results ====")
	fmt.Printf("Earnings before Tax: %.2f\n", ebt)
	fmt.Printf("Profit after Tax: %.2f\n", profit)
	fmt.Printf("Ratio Between EBT and Profit: %.2f\n", ratio)
	fmt.Printf("\n")
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