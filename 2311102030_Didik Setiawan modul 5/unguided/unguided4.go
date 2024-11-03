package main

import (
	"fmt"
)

func printSeries(number, current int) {
	if current == 1 {
		fmt.Printf("%d ", current)
		return
	}

	fmt.Printf("%d ", current)
	printSeries(number, current-1)
	fmt.Printf("%d ", current)
}

func main() {
	var number int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&number)

	fmt.Printf("Barisan bilangan dari %d hingga 1 dan kembali ke %d:\n", number, number)
	printSeries(number, number)
	fmt.Println()
}
