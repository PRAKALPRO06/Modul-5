package main

import "fmt"

// fungsi untuk mencetak bilangan dari n hingga 1
func cetakMundur(n int) {
	if n <= 0 {
		return
	}

	fmt.Println(n)
	cetakMundur(n - 1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n untuk mencetak bilangan dari n hingga 1: ")
	fmt.Scanln(&n)
	fmt.Println("Hasil cetak mundur:")
	cetakMundur(n)
}
