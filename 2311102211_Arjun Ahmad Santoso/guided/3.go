package main 

import "fmt"

func pangkatDua(n int) int {
	if n == 0 {
		return 1
	}
	return 2 * pangkatDua(n-1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n untuk mencari 2 pangkat n: ")
	fmt.Scan(&n)
	fmt.Println("hasil dari 2 pangkat", n, "adalah", pangkatDua(n))
}
