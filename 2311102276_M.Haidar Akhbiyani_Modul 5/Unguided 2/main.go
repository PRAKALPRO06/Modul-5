package main

import "fmt"

func PolaAsterisk(n int) {
	if n == 0 {
		return
	}
	PolaAsterisk(n - 1)

	for i := 0; i < n; i++ {
		fmt.Print("*")
	}

	fmt.Println()
}

func main() {

	var n int

	fmt.Print("Masukkan nilai n : ")
	fmt.Scan(&n)

	PolaAsterisk(n)
}
