package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {

	var username string
	var currOption int
	loop := true

	for loop {
		homePage()
		fmt.Scan(&currOption)

		if currOption == 1 {
			fmt.Println("Your name?")
			fmt.Scan(&username)
			openNewAccount(username)
		} else if currOption == 2 {
			if username == "" {
				fmt.Println("Your name?")
				fmt.Scan(&username)
			}

			account := isAccountExists(username)
			if !account {
				fmt.Println("Account does not exist. Do you want to create a new Account?")
				userCreatesAnewAccount := false
				fmt.Scan(&userCreatesAnewAccount)
				if userCreatesAnewAccount {
					openNewAccount(username)
				}
			}

			balance := getBalance(username)
			fmt.Printf("Your account balance is %f0.2", balance)

		} else if currOption == 3 {
			if username == "" {
				fmt.Println("Your name?")
				fmt.Scan(&username)
			}

			account := isAccountExists(username)
			if !account {
				fmt.Println("Account does not exist. Do you want to create a new Account?")
				userCreatesAnewAccount := false
				fmt.Scan(&userCreatesAnewAccount)
				if userCreatesAnewAccount {
					openNewAccount(username)
				}
			}

			fmt.Println("Enter deposit amount: ")
			var amount float64
			fmt.Scan(&amount)
			depositAmount(username, amount)

		} else if currOption == 4 {
			if username == "" {
				fmt.Println("Your name?")
				fmt.Scan(&username)
			}

			account := isAccountExists(username)
			if !account {
				fmt.Println("Account does not exist. Do you want to create a new Account?")
				userCreatesAnewAccount := false
				fmt.Scan(&userCreatesAnewAccount)
				if userCreatesAnewAccount {
					openNewAccount(username)
				}
			}

			fmt.Println("Enter withdraw amount: ")
			var amount float64
			fmt.Scan(&amount)
			withdrawAmount(username, amount)

		} else if currOption == 5 {
			loop = false
		}
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

	account := isAccountExists(username)
	if !account {
		subfolder := "accounts"
		filename := username + ".txt"
		filePath := filepath.Join(subfolder, filename)

		os.MkdirAll(subfolder, os.ModePerm)
		os.WriteFile(filePath, []byte("0"), 0644)
		fmt.Println("Your account has been created!")
		return
	}

	fmt.Println("You already have an account")
}

func isAccountExists(username string) bool {
	subfolder := "accounts"
	filename := username + ".txt"
	filePath := filepath.Join(subfolder, filename)

	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func getBalance(username string) float64 {
	subfolder := "accounts"
	filename := username + ".txt"
	filePath := filepath.Join(subfolder, filename)

	data, _ := os.ReadFile(filePath)

	balance, _ := strconv.ParseFloat(string(data), 64)

	return balance
}

func depositAmount(username string, deposit float64) {
	currentBalance := getBalance(username)
	updatedBalance := strconv.FormatFloat(currentBalance+deposit, 'f', -1, 64)

	subfolder := "accounts"
	filename := username + ".txt"
	filePath := filepath.Join(subfolder, filename)
	os.WriteFile(filePath, []byte(updatedBalance), 0644)

	fmt.Printf("Success! Your new balance is: %v0.1 \n", updatedBalance)
}

func withdrawAmount(username string, withdraw float64) {
	currentBalance := getBalance(username)
	updatedBalance := strconv.FormatFloat(currentBalance-withdraw, 'f', -1, 64)

	subfolder := "accounts"
	filename := username + ".txt"
	filePath := filepath.Join(subfolder, filename)
	os.WriteFile(filePath, []byte(updatedBalance), 0644)

	fmt.Printf("Success! Your new balance is: %v0.1 \n", updatedBalance)
}
