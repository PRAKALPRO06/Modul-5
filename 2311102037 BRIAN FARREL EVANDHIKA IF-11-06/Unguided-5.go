//2311102037 BRIAN FARREL EVANDHIKA IF 11 06
package main

import "fmt"

// Fungsi rekursif untuk mencetak bilangan ganjil dari 1 hingga N
func printOdd(n, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	printOdd(n, current+2) // Melompat ke bilangan ganjil berikutnya
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)

	fmt.Print("Bilangan ganjil dari 1 hingga ", n, ": ")
	printOdd(n, 1) // Mulai dari 1
	fmt.Println()
}
