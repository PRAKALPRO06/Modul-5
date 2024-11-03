package main

import "fmt"

func faktorBilangan(n, hasil int) {
	if hasil > n {
		return
	}
	if n%hasil == 0 {
		fmt.Print(hasil, " ")
	}
	faktorBilangan(n, hasil+1)
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat : ")
	fmt.Scan(&n)
	fmt.Print("Faktor dari ", n, ": ")
	faktorBilangan(n, 1)
}
