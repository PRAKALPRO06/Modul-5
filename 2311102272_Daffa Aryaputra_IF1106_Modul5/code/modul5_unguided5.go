package main

import (
	"fmt"
)

func printOddSeries(n, current int) {
	if current <= n {
		fmt.Print(current, " ")
		printOddSeries(n, current+2)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	printOddSeries(n, 1)
	fmt.Println()
}
