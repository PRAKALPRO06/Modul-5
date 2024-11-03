package main

import "fmt"

func bintang(hitung int) {
	if hitung > 0 {
		fmt.Print("*")
		bintang(hitung - 1)
	}
}

func cetak(n int, baris int) {
	if baris <= n {
		bintang(baris)
		fmt.Println()
		cetak(n, baris + 1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)
	cetak(n, 1)
}
