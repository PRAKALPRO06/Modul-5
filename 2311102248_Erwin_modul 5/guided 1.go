package main

import "fmt"

func cetakMundur(n int) {
	if n == 1 {
		fmt.Println(n)
		return
	}
	fmt.Print(n, " ")
	cetakMundur(n - 1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n untuk cetak bilangan dari n hingga 1: ")
	fmt.Scanln(&n)
	fmt.Print("Hasil cetak mundur: ")
	cetakMundur(n)
}
