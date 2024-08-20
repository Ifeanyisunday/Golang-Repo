package main

import "fmt"

func main() {
	var miles float64
	var gallon float64

	fmt.Println("Enter miles")
	fmt.Scanln(&miles)
	fmt.Println("Enter gallon")
	fmt.Scanln(&gallon)

	var result float64 = gallon / miles
	fmt.Printf("%s %.2f", "gallon per mile =", result)
}
