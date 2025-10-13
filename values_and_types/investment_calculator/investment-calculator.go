package main

import (
	"fmt"
	"math"
)

func main() {

	//one way to declare variables
	var initialValue int
	var expectedReturnRate float64 = 5.5
	var years float64
	const inflationRate = 2.5 //this is a constant, so the value cannot be modified

	banner := `
	 ___                     _                        _   
	|_ _|_ ____   _____  ___| |_ _ __ ___   ___ _ __ | |_ 
	 | || '_ \ \ / / _ \/ __| __| '_  _ \ / _ \ '_ \| __|
	 | || | | \ V /  __/\__ \ |_| | | | | |  __/ | | | |_ 
	|___|_| |_|\_/ \___||___/\__|_| |_| |_|\___|_| |_|\__|
 / ___|__ _| | ___ _   _| | __ _| |_ ___  _ __        
| |    / _ | |/ __| | | | |/ _ | __/ _ \| __|       
| |__| (_| | | (__| |_| | | (_| | || (_) | |          
 \____\__,_|_|\___|\__,_|_|\__,_|\__\___/|_|          
	`

	fmt.Println(banner)
	fmt.Print("Input initial investment value: ")
	fmt.Scan(&initialValue) //to ask user input for scanning, we need to pass a pointer to that value (using &)

	fmt.Print("In how many years: ")
	fmt.Scan(&years)

	//a second way, go will automatically declare the type for us (this is generally a better way)
	futureValue := float64(initialValue) * math.Pow((1+expectedReturnRate/100), years)
	realFutureValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Println()
	fmt.Println("===== RESULTS =====")
	fmt.Println("Without inflation: ", futureValue)
	fmt.Println("With 2.5 percent inflation: ", realFutureValue)
}
