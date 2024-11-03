package main

import "fmt"

// Fungsi untuk mencetak spasi pada baris tertentu
func cetakSpasi(n int) {
	if n == 0 {
		return
	}
	fmt.Print(" ")
	cetakSpasi(n - 1)
}

// Fungsi untuk mencetak bintang pada baris tertentu
func cetakBintang(i int) {
	if i == 0 {
		return
	}
	fmt.Print("*")
	cetakBintang(i - 1)
}

// Fungsi untuk mencetak pola bintang
func cetakPola(n, i int) {
	if i > n {
		return // Basis rekursif: jika sudah mencetak semua baris
	}

	// Mencetak spasi untuk membentuk segitiga
	cetakSpasi(n - i)

	// Mencetak bintang pada baris ke-i
	cetakBintang(2*i - 1)

	fmt.Println() // Pindah ke baris berikutnya

	// Rekursi untuk mencetak baris berikutnya
	cetakPola(n, i+1)
}

func main() {
	var n int

	fmt.Print("Masukkan jumlah baris pola bintang: ")
	fmt.Scanln(&n)

	cetakPola(n, 1) // Memanggil fungsi untuk mencetak pola bintang
}
