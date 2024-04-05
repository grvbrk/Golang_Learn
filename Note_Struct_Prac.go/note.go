package main

import (
	"errors"
	"fmt"
)

func main() {

}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("No input provided")
	}

	return value, nil
}
