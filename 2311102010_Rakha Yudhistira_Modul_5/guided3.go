//Rakha Yudhistira_2311102010_IF-11-06
package main

import "fmt"

// Fungsi untuk mencari 2 pangkal n
func pangkatDua(n int) int {
	if n == 0 {
		return 1
	}
	return 2 * pangkatDua(n-1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n untuk mencari 2 pangkat n: ")
	fmt.Scanln(&n)
	fmt.Println("Hasil 2 pangkat", n, ":", pangkatDua(n))
}