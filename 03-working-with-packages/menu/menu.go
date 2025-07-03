package menu

import "fmt"

func PrintTitleBar() {
	fmt.Println("=========================")
	fmt.Println("   Welcome to Go Bank!")
	fmt.Println("=========================")
}

func PrintMenu() {
	fmt.Println("\n[?] What do you want to do?:")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}
