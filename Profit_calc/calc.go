package main

import "fmt"

func main(){
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter your profit: ")
	fmt.Scan( &expenses)
	fmt.Print("Enter your tax rate: ")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses
	var profit = EBT*(1-taxRate/100)
	var ratio = EBT/profit

	fmt.Println("EBT", EBT)
	fmt.Println("Profit", profit)
	fmt.Println("Ratio", ratio)

}