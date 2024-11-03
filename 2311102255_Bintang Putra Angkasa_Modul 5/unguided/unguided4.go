package main

import "fmt"

// Fungsi rekursif untuk menampilkan bilangan dari N hingga 1, lalu kembali ke N
func tampilBarisan(N, i int) {
	if i == 1 {
		fmt.Print(i, " ") // Mencetak 1 ketika mencapai titik tengah
		return
	}

	fmt.Print(i, " ")     // Mencetak bilangan dari N hingga 2
	tampilBarisan(N, i-1) // Rekursi menuju ke bawah hingga 1
	fmt.Print(i, " ")     // Mencetak bilangan kembali ke N
}

func main() {
	var N int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&N)

	fmt.Print("Barisan: ")
	tampilBarisan(N, N) // Memanggil fungsi untuk menampilkan barisan
	fmt.Println()
}
