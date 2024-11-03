// 2311102012
package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung x pangkat y
func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan basis (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan pangkat (y): ")
	fmt.Scan(&y)

	hasil := pangkat(x, y)
	fmt.Printf("Hasil %d pangkat %d adalah: %d\n", x, y, hasil)
}
