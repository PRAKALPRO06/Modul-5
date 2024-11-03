package main

import "fmt"

func cetak(n int) {
	if n == 1 {
		fmt.Println(n)
		return
	}
	fmt.Print(n, "")
	cetak(n - 1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n untuk cetak bilangan n hingga 1 : ")
	fmt.Scanln(&n)
	fmt.Print("Hasil cetak mundur : ")
	cetak(n)
}
