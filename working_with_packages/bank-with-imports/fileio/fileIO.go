package fileio

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func WriteToFile(balance float64) {
	balanceText := fmt.Sprint("Current balance: ", balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func GetBalanceFromFile() (float64, error) {
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
