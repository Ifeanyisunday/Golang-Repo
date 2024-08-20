package main

import "fmt"

func main() {
	var name string
	var salary int

	fmt.Println("Enter citizen name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter citizen year salary: ")
	fmt.Scanln(&salary)

	var getTax float64
	var taxPer1 int
	var taxPer2 int

	if salary <= 30000 {
		taxPer1 = 15 / 100
		getTax = float64(salary * taxPer1)
		fmt.Println("your tax is", getTax)
	} else {
		taxPer2 = 20 / 100
		getTax = float64(salary * taxPer2)
		fmt.Println("your tax is", getTax)
	}

}
