//Rakha Yudhistira_2311102010_IF-11-06
package main

import "fmt"

// Fungsi untuk menghitung faktorial n!
func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n untuk mencari faktorial n!: ")
	fmt.Scanln(&n)
	fmt.Println("Hasil faktorial dari", n, ":", faktorial(n))
}