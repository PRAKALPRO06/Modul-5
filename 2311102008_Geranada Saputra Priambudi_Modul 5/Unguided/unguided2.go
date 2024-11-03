package main

import "fmt"

func polaBintang(n int, baris int) {
	if baris > n {
		return
	}
	for i := 0; i < baris; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	polaBintang(n, baris+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	fmt.Println("Pola bintang:")
	polaBintang(n, 1)
}
