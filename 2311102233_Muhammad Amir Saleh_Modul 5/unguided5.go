package main

import "fmt"

func ganjil(n, angka int) {
	if angka > n {
		return
	}
	fmt.Print(angka, " ")
	ganjil(n, angka+2)
}

func main() {
	var n int
	fmt.Scan(&n)
	ganjil(n, 1)
}
