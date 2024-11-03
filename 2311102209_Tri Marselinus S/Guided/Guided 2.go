package main

import "fmt"

func jumlahRekursif ( n int) int {
	if n==1 {
		return 1
	}
	return n + jumlahRekursif(n-1)
}

func main (){
	var n int
	fmt.Print("Masukkan nilai n untuk penjumlahan dari n hingga 1: ")
	fmt.Scanln(&n)
	fmt.Print("Hasil cetak penjumlahan:",jumlahRekursif(n))
}