package main

import "fmt"


// Fungsi untuk menghitung faktorial dari n!
func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}

func main() {
	var n int
	fmt.Println("Masukan nilai n untuk mencari faktorial n!: ")
	fmt.Scan(&n)
	fmt.Println("Hasil faktorial dari", n, ":", faktorial(n))
}
