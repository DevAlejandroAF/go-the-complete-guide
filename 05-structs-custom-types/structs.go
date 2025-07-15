package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs/note"
	//"example.com/structs/user"
)

// Note code
func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("[!] Saving the note failed.")
		return
	}

	fmt.Println("[+] Note saved successfully!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

// User code
// func main() {
// 	fmt.Println("Structs App")
// 	fmt.Println()
//
// 	userFirstName := user.GetUserData("[+] Please enter your first name: ")
// 	userLastName := user.GetUserData("[+] Please enter your last name: ")
// 	userBirthDate := user.GetUserData("[+] Please enter your birthdate (MM/DD/YYYY): ")
//
// 	var appUser *user.User
// 	appUser, err := user.New(userFirstName, userLastName, userBirthDate)
//
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
//
// 	appUser.OutputUserDetails()
// 	appUser.ClearUserName()
//
// 	adminFirstName := user.GetUserData("[+] Please enter your first name: ")
// 	adminLastName := user.GetUserData("[+] Please enter your last name: ")
// 	adminBirthDate := user.GetUserData("[+] Please enter your birthdate (MM/DD/YYYY): ")
// 	adminEmail := user.GetUserData("[+] Please enter your email: ")
// 	adminPassword := user.GetUserData("[+] Please enter your password: ")
//
// 	var appAdmin *user.Admin
// 	appAdmin, err = user.NewAdmin(adminFirstName, adminLastName, adminBirthDate, adminEmail, adminPassword)
//
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
//
// 	appAdmin.OutputAdminDetails()
// 	appAdmin.ClearUserName()
// }
