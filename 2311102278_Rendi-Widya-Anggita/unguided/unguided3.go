package main

import "fmt"
func faktor(n int, bil int) {
	if bil > n {
		return
	}
	if n%bil == 0 {
		fmt.Print(bil, " ")
	}
	faktor(n, bil+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	fmt.Print("Faktor dari ", n, ": ")
	faktor(n, 1)
	fmt.Println()
}
