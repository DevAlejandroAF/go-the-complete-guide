package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Failed to read the file.")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to convert value.")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Printf("[!] Sorry there was an error reading the %v file\n", accountBalanceFile)
		fmt.Printf("[!] Error: %v\n", err)
		fmt.Println("[!] Setting balance to a test value.")
	}
	fmt.Println("=========================")
	fmt.Println("   Welcome to Go Bank!")
	fmt.Println("=========================")

	for {
		fmt.Println("[?] What do you want to do?:")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		choiceInput, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(choiceInput))

		switch choice {
		case 1:
			fmt.Printf("[+] Your balance is: %v\n", accountBalance)
		case 2:
			fmt.Println("[+] Type your deposit amount: ")
			depositInput, _ := reader.ReadString('\n')
			deposit, _ := strconv.ParseFloat(strings.TrimSpace(depositInput), 64)
			if deposit <= 0 {
				fmt.Println("[!] Your deposit can't be zero or a negative value")
			} else {
				accountBalance += deposit
				fmt.Printf("[+] Your new balance is: %v\n", accountBalance)
				writeBalanceToFile(accountBalance)
			}
		case 3:
			fmt.Println("[?] How much money you want to Withdraw: ")
			withdrawInput, _ := reader.ReadString('\n')
			withdraw, _ := strconv.ParseFloat(strings.TrimSpace(withdrawInput), 64)
			if withdraw > accountBalance {
				fmt.Println("[!] Your withdraw can't be superior than your account money")
			} else if withdraw <= 0 {
				fmt.Println("[!] Your withdraw can't be zero or a negative value")
			} else {
				accountBalance -= withdraw
				fmt.Printf("[+] Your withdraw for %v was successful, your new balance is %v\n", withdraw, accountBalance)
				writeBalanceToFile(accountBalance)
			}
		case 4:
			fmt.Println("[+] Thanks for choosing us, goodbye!")
			return
		}
	}
}
