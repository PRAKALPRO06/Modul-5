package main

import "fmt"

func printStars(n int, current int) {
	// Basis case - berhenti jika current melebihi n
	if current > n {
		return
	}

	// Print bintang untuk baris saat ini
	for i := 0; i < current; i++ {
		fmt.Print("*")
	}
	fmt.Println()

	// Panggil rekursif untuk baris selanjutnya
	printStars(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	// Mulai dari 1 bintang
	printStars(n, 1)
}
