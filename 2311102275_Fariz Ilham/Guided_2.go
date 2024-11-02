package main

import "fmt"

// fungsi untuk menghitung penjumlahan 1 hingga n
func jumlahRekursif(n int) int {
	if n == 1 {
		return 1
	}
	return n + jumlahRekursif(n-1)
}

func main() {
	var n int
	fmt.Print(" Masukkan nilai n untuk penjumlahan 1 hingga n: ")
	fmt.Scanln(&n)
	fmt.Println(" hasil penjumlahan: ", jumlahRekursif(n))
}
