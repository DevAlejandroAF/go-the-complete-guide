package main

import (
	"bufio"
	"example.com/packages/fileops"
	"example.com/packages/menu"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var accountBalanceFile = "balance.txt"

func main() {
	reader := bufio.NewReader(os.Stdin)
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile, 1000)
	if err != nil {
		fmt.Printf("[!] Sorry there was an error reading the %v file\n", accountBalanceFile)
		fmt.Printf("[!] Error: %v\n", err)
		fmt.Println("[!] Setting balance to a test value.")
	}
	menu.PrintTitleBar()

	for {
		menu.PrintMenu()

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
				fileops.WriteOutputToFile(accountBalance, accountBalanceFile)
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
				fileops.WriteOutputToFile(accountBalance, accountBalanceFile)
			}
		case 4:
			fmt.Println("[+] Thanks for choosing us, goodbye!")
			return
		}
	}
}
