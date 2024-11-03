package main

import "fmt"

func main() {
	var c int
	var b string

	fmt.Print("Masukkan jumlah baris pola: ")
	fmt.Scan(&c)

	pola_bintang(1, c, b)
}

func pola_bintang(i, c int, b string) {

	if i <= c {

		b = "*" + b

		fmt.Println(b)

		pola_bintang(i+1, c, b)
	}
}
