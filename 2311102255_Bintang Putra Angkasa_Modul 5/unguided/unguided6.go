package main

import "fmt"

// Fungsi rekursif untuk menghitung x pangkat y
func pangkat(x, y int) int {
	if y == 0 {
		return 1 // Basis rekursif: jika pangkat y adalah 0, hasilnya adalah 1
	}
	return x * pangkat(x, y-1) // Rekursi dengan mengurangi y hingga 0
}

func main() {
	var x, y int

	fmt.Print("Masukkan bilangan x: ")
	fmt.Scanln(&x)

	fmt.Print("Masukkan bilangan y: ")
	fmt.Scanln(&y)

	hasil := pangkat(x, y) // Memanggil fungsi rekursif untuk menghitung x^y
	fmt.Printf("Hasil %d pangkat %d adalah: %d\n", x, y, hasil)
}
