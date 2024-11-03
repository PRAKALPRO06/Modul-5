package main

import "fmt"

func jumlahRekursi(n int) int {
	if n == 1 {
		return 1
	}
	return n + jumlahRekursi(n-1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n untuk penjumlahan 1 hingga n: ")
	fmt.Scanln(&n)
	fmt.Println("Hasil penjumlahan:", jumlahRekursi(n))
}
