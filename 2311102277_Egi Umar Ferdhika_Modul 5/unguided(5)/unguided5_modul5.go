package main

import "fmt"

func printOddNumbers(current, n int) {
	// Basis case - berhenti jika current melebihi n
	if current > n {
		return
	}

	// Cetak bilangan jika ganjil
	if current%2 != 0 {
		fmt.Printf("%d ", current)
	}

	// Rekursi ke bilangan berikutnya
	printOddNumbers(current+1, n)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Mohon masukkan bilangan positif!")
		return
	}

	fmt.Print("Bilangan ganjil dari 1 hingga ", n, ": ")
	printOddNumbers(1, n)
	fmt.Println()
}
