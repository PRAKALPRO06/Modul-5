// 2311102012
package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak bintang dalam satu baris
func cetakBintang(jumlah int) {
	if jumlah > 0 {
		fmt.Print("*")
		cetakBintang(jumlah - 1)
	}
}

// Fungsi rekursif untuk mencetak pola bintang dari 1 hingga jumlah
func cetakPola(jumlah, current int) {
	if current <= jumlah {
		cetakBintang(current)
		fmt.Println()
		cetakPola(jumlah, current+1)
	}
}

func main() {
	var jumlah int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&jumlah)

	fmt.Println("Pola bintang:")
	cetakPola(jumlah, 1)
}
