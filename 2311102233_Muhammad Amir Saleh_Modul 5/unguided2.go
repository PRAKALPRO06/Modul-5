package main

import "fmt"

func cetakBintang(n, baris, kolom int) {
	if baris > n {
		return
	} else if kolom <= baris {
		fmt.Print("*")
		cetakBintang(n, baris, kolom+1)
		return
	}

	fmt.Println()
	cetakBintang(n, baris+1, 1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	cetakBintang(n, 1, 1)
}
