package main

import "fmt"

func power(x, y int) int {
	// Basis case
	if y == 0 {
		return 1
	}
	if y == 1 {
		return x
	}

	// Rekursi: x * x^(y-1)
	return x * power(x, y-1)
}

func main() {
	var x, y int

	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan pangkat y: ")
	fmt.Scan(&y)

	if y < 0 {
		fmt.Println("Mohon masukkan pangkat non-negatif!")
		return
	}

	hasil := power(x, y)
	fmt.Printf("Hasil %d pangkat %d adalah: %d\n", x, y, hasil)
}
