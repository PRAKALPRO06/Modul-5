// 2311102012
package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak faktor bilangan
func cetakFaktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	cetakFaktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	fmt.Printf("Faktor dari %d adalah: ", n)
	cetakFaktor(n, 1)
	fmt.Println()
}
