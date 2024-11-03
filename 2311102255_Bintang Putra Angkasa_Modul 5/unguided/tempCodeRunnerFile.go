package main

import "fmt"

// Fungsi rekursif untuk menampilkan faktor-faktor dari N
func tampilkanFaktor(N, i int) {
	if i > N {
		return // Basis rekursif: jika i sudah lebih besar dari N, fungsi berhenti
	}

	if N%i == 0 {
		fmt.Print(i, " ") // Mencetak i jika i adalah faktor dari N
	}

	tampilkanFaktor(N, i+1) // Rekursi untuk mengecek faktor berikutnya
}

func main() {
	var N int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&N)

	fmt.Print("Faktor dari ", N, " adalah: ")
	tampilkanFaktor(N, 1) // Memanggil fungsi untuk menampilkan faktor dari N, mulai dari 1
	fmt.Println() // Baris baru setelah semua faktor ditampilkan
}
