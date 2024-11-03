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
	fmt.Print("Masukkan basis (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan pangkat (y): ")
	fmt.Scan(&y)

	hasil := pangkat(x, y)
	fmt.Printf("Hasil dari %d^%d adalah: %d\n", x, y, hasil)
}
