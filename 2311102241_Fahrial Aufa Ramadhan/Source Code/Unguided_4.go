package main

import (
	"fmt"
)

func urutan(n, current int) {
	if current < 1 {
		return
	}

	fmt.Print(current, " ")

	if current > 1 {
		urutan(n, current-1)
	}

	fmt.Print(current, " ")
}

func main() {
	var n int

	fmt.Print("Masukkan nilai : ")
	fmt.Scan(&n)

	urutan(n, n)
	fmt.Println()
}
