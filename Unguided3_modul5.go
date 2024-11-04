package main

import "fmt"

func faktorBilangan(x, M int) {
	if M == x {
		fmt.Print(M)
	} else {
		if x%M == 0 {
			fmt.Print(M, " ")
			faktorBilangan(x, M+1)
		} else {
			faktorBilangan(x, M+1)
		}

	}
}

func main() {
	var x int
	fmt.Print("Masukkan angka : ")
	fmt.Scan(&x)
	faktorBilangan(x, 1)
}
