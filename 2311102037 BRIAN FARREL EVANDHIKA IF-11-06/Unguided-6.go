//2311102037 BRIAN FARREL EVANDHIKA IF 11 06
package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung x^y
func power(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * power(x, y-1)
}

func main() {
	var x, y int

	// Meminta input dari user
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan y: ")
	fmt.Scan(&y)

	// Menghitung hasil pangkat dan menampilkan
	result := power(x, y)
	fmt.Printf("Hasil %d pangkat %d adalah %d\n", x, y, result)
}
