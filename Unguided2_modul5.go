package main

import (
	"fmt"
)

func cetakBaris(nomorBintang int) {
	if nomorBintang == 0 {
		return
	}
	fmt.Print("*")
	cetakBaris(nomorBintang - 1)
}

func cetakBintang(x int, current int) {
	if current > x {
		return
	}

	cetakBaris(current)
	fmt.Println()

	cetakBintang(x, current+1)
}

func main() {
	var x int
	fmt.Print("Masukkan nomor : ")
	fmt.Scan(&x)

	cetakBintang(x, 1)
}
