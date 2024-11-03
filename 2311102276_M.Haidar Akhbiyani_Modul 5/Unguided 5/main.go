package main

import "fmt"

func Ganjil(n, hasil int) {
	if hasil > n {
		return
	}
	if hasil%2 != 0 {
		fmt.Print(hasil, " ")
	}

	Ganjil(n, hasil+1)
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat n : ")
	fmt.Scan(&n)
	fmt.Print("Hasil bilangan ganjil dari ", n, " : ")
	Ganjil(n, 1)
}
