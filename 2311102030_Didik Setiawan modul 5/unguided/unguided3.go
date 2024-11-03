package main

import (
	"fmt"
)

func printFactors(number, divisor int) {
	if divisor > number {
		return
	}

	if number%divisor == 0 {
		fmt.Printf("%d ", divisor)
	}

	printFactors(number, divisor+1)
}

func main() {
	var number int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&number)

	fmt.Printf("Faktor-faktor dari %d adalah: ", number)
	printFactors(number, 1)
	fmt.Println()
}
