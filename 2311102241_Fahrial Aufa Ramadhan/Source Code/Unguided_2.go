package main

import "fmt"

func bintang(n int) {
	if n > 0 {
		fmt.Print("*")
		bintang(n - 1)
	}
}

func pola(n int) {
	if n > 0 {
		pola(n - 1)
		bintang(n)
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)

	fmt.Println("Pola bintang:")
	pola(n)
}
