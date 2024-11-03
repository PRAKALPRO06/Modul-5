//2311102018_ Aryo Tegar Sukarno
package main

import "fmt"

func cetakmundur(n int) {
	if n < 1 {
		return // Menghentikan rekursi jika n kurang dari 1
	}
	fmt.Print(n, " ") // Menggunakan Print untuk mencetak di baris yang sama
	cetakmundur(n - 1)
}

func main() {
	var n int
	fmt.Print("Masukkan Nilai N: ")
	fmt.Scanln(&n)
	fmt.Print("Hasil dari cetak mundur: ")
	cetakmundur(n)
	fmt.Println() // Menambahkan baris baru setelah output
}