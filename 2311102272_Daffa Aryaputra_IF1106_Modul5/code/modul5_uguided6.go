package main

import (
	"fmt"
)

func power(x, y int) int {
	if y == 0 {
		return 1
	}

	return x * power(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	result := power(x, y)
	fmt.Printf("%d pangkat %d adalah %d\n", x, y, result)
}
