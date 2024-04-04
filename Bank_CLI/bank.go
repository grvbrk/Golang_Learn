package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	var username string
	var currOption int

	homePage()
	fmt.Scan(&currOption)

	if currOption == 1 {
		fmt.Println("Your name?")
		fmt.Scan(&username)
		openNewAccount(username)
	} else if currOption == 2 {

	} else if currOption == 3 {

	} else if currOption == 4 {

	} else if currOption == 5 {
		return
	}
}

func homePage() {
	fmt.Println("Welcome to the bank! Please choose an option")
	fmt.Println("1. Open an Account")
	fmt.Println("2. Check Balance")
	fmt.Println("3. Deposit Money")
	fmt.Println("4. Withdraw Money")
	fmt.Println("5. Exit App")
}

func openNewAccount(username string) {
	subfolder := "accounts"
	filename := username + ".txt"
	filePath := filepath.Join(subfolder, filename)

	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(subfolder, os.ModePerm)
			os.WriteFile(filePath, []byte("0"), 0644)
			return
		}
	}

	fmt.Println("You already have an account")

}
