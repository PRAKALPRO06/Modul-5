// 2311102012
package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak bilangan ganjil dari 1 hingga n
func cetakGanjil(n, i int) {
	if i > n {
		return
	}
	if i%2 != 0 {
		fmt.Print(i, " ")
	}
	cetakGanjil(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	fmt.Printf("Bilangan ganjil dari 1 hingga %d adalah: ", n)
	cetakGanjil(n, 1)
	fmt.Println()
}
