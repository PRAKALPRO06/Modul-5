//2311102037 BRIAN FARREL EVANDHIKA IF 11 06
package main

import "fmt"

// Fungsi rekursif untuk mencetak barisan menurun dari N ke 1
func printDown(n int) {
	if n < 1 {
		return
	}
	fmt.Print(n, " ")
	printDown(n - 1)
}

// Fungsi rekursif untuk mencetak barisan naik dari 1 ke N
func printUp(n, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	printUp(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)

	fmt.Print("Hasil: ")
	printDown(n)
	printUp(n, 2) // Mulai dari 2 agar tidak mencetak 1 dua kali
	fmt.Println()
}
