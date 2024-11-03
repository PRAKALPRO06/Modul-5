package main

import "fmt"

func tampil(n int) {
	if n > 0 {
		fmt.Print(n, " ")
		tampil(n - 1)
	}
}
func main() {
	var n int
	fmt.Print("Masukkan bilangannya: ")
	fmt.Scan(&n)
	tampil(n)
	fmt.Println()
}
