package main

import (
	"fmt"
)

func cetakBilangan(n int) {
	if n > 0 {
		fmt.Print(n, " ")
		cetakBilangan(n - 1)
		fmt.Print(n, " ")
	}
}

func main() {
	var N int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	cetakBilangan(N)
	fmt.Println()
}
