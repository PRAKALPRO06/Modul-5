package main

import "fmt"

// Fungsi rekursif untuk menampilkan bilangan ganjil dari 1 hingga N
func tampilBilanganGanjil(i, N int) {
	if i > N {
		return // Basis rekursif: jika i lebih besar dari N, fungsi berhenti
	}

	if i%2 != 0 {
		fmt.Print(i, " ") // Mencetak i jika i adalah bilangan ganjil
	}

	tampilBilanganGanjil(i+1, N) // Rekursi untuk memeriksa bilangan berikutnya
}

func main() {
	var N int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&N)

	fmt.Print("Barisan bilangan ganjil dari 1 hingga ", N, " adalah: ")
	tampilBilanganGanjil(1, N) // Memanggil fungsi untuk menampilkan bilangan ganjil dari 1 hingga N
	fmt.Println()
}
