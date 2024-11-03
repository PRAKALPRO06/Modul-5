package main

import (
	"fmt"
)

func printtambahangka(n int, current int) {
	if current > n {
		return
	}
	if current%2 != 0 {
		fmt.Printf("%d ", current)
	}
	printtambahangka(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	fmt.Print("Keluaran: ")
	printtambahangka(n, 1)
	fmt.Println()
}
