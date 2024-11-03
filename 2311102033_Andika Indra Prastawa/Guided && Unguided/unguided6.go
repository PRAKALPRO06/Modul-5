package main

import (
	"fmt"
)

func pangkat(basis int, eksponen int) int {

	if eksponen == 0 {
		return 1
	}

	return basis * pangkat(basis, eksponen-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan y: ")
	fmt.Scan(&y)

	hasil := pangkat(x, y)
	fmt.Printf("Keluaran: %d\n", hasil)
}
