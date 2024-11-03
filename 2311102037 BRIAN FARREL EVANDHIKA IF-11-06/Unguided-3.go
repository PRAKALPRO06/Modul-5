//2311102037 BRIAN FARREL EVANDHIKA IF 11 06
package main

import "fmt"

// Fungsi rekursif untuk menemukan faktor
func findFactors(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	findFactors(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)

	fmt.Print("Faktor dari ", n, ": ")
	findFactors(n, 1)
	fmt.Println()
}
