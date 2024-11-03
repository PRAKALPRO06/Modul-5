package main

import "fmt"

func pangkat(x int, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan y: ")
	fmt.Scan(&y)

	hasil := pangkat(x, y)
	fmt.Printf("Hasil %d pangkat %d adalah: %d\n", x, y, hasil)
}
