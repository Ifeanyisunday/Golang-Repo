package main

import "fmt"

func main() {

	var acctNo int
	var beginingBalance int
	var totalItem int
	var totalCredit int
	var creditLimit int

	fmt.Println("Enter Account No: ")
	fmt.Scanln(&acctNo)
	fmt.Println("Enter beginning balance: ")
	fmt.Scanln(&beginingBalance)
	fmt.Println("Enter total item: ")
	fmt.Scanln(&totalItem)
	fmt.Println("Enter total credit: ")
	fmt.Scanln(&totalCredit)
	fmt.Println("Enter credit limit")
	fmt.Scanln(&creditLimit)

	var newBalance int = beginingBalance + totalItem - totalCredit
	fmt.Println("your new balance: ", newBalance)

	if newBalance > creditLimit {
		fmt.Println("Credit limit exceeded")
	} else {
		fmt.Println("you are within the credit limit")
	}
}
