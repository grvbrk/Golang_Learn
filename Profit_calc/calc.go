package main

import "fmt"

func main(){
	var revenue float64
	var expenses float64
	var taxRate float64

	getUserInput("Enter your revenue: ", &revenue)
	getUserInput("Enter your profit: ", &expenses)
	getUserInput("Enter your tax rate: ", &taxRate)

	EBT := revenue - expenses
	var profit = EBT*(1-taxRate/100)
	var ratio = EBT/profit

	fmt.Println("EBT", EBT)
	fmt.Println("Profit", profit)
	fmt.Println("Ratio", ratio)

}

func getUserInput(text string, variable *float64){
	fmt.Print(text)
	fmt.Scan(variable)
}