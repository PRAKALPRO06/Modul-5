package main

import (
	"fmt"
)

func printOddNumbers(limit, current int) {
	if current > limit {
		return
	}
	if current%2 != 0 {
		fmt.Printf("%d ", current)
	}
	printOddNumbers(limit, current+1)
}

func main() {
	var limit int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&limit)

	fmt.Printf("Bilangan ganjil dari 1 hingga %d adalah:\n", limit)
	printOddNumbers(limit, 1)
	fmt.Println()
}
