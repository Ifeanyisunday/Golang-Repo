package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("N\t\t", "N2\t\t", "N3\t\t", "N4")
	for i := 1.0; i <= 5.0; i++ {
		for j := 1.0; j <= 4.0; j++ {
			fmt.Printf("%.0f\t    ", math.Pow(i, j))
		}
		fmt.Println()
	}
}
