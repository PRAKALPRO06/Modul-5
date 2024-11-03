package main

import (
	"fmt"
)

func tampilkanGanjil(n int, i int) {
	if i > n {
		return
	}
	if i%2 != 0 {
		fmt.Print(i, " ")
	}
	tampilkanGanjil(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	fmt.Print("Keluaran: ")
	tampilkanGanjil(n, 1)
	fmt.Println()
}
