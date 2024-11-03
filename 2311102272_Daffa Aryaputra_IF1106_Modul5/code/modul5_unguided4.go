package main

import (
	"fmt"
)

func printSequence(n, current int) {
	if current < 1 {
		return
	}

	fmt.Print(current, " ")

	if current > 1 {
		printSequence(n, current-1)
	}

	fmt.Print(current, " ")
}

func main() {
	var n int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	printSequence(n, n)
	fmt.Println()
}