package main

import "fmt"

func cetakFaktor(n int, current int) {
	if current > n {
		return
	}

	if n%current == 0 {
		fmt.Print(current)
		if current != n {
			fmt.Print(" ")
		}
	}
	cetakFaktor(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Mohon masukkan bilangan positif!")
		return
	}

	fmt.Print("Faktor Bilangan dari ", n, " adalah : ")
	cetakFaktor(n, 1)
	fmt.Println()
}
