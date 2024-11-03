package main

import "fmt"

// fungsi untuk menghitung penjumlahan dari 1 hingga n
func jumlahRekursif(n int) int {
	if n == 1 {
		return 1
	} 
	return n + jumlahRekursif(n-1)
}
func main() {
	var n int
	fmt.Print("Masukan nilai n untuk penjumlahan dari 1 hingga n: ")
	fmt.Scan(&n)
	fmt.Println("Hasil penjumlahan: ", jumlahRekursif(n))
}