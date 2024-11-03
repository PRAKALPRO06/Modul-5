package main

import "fmt"

// Fungsi rekursif untuk menghitung 2^n
func pangkatDua(n int) int {
	if n == 0 {
		return 1
	} 
	return 2 * pangkatDua(n-1)
}

func main() {
	var n int
	fmt.Println("Masukan nilai n untuk mencari 2 pangkat n: ")
	fmt.Scan(&n)
	fmt.Println("Hasil 2 pangkat", n, ":", pangkatDua(n))

}
