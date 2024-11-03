package main

import "fmt"

func baris(n, hasil int) {
	if n == hasil {
		fmt.Print(1, " ")
		return
	}
	fmt.Print(n, " ")
	baris(n-1, hasil)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat n : ")
	fmt.Scan(&n)
	fmt.Print("Hasil : ")
	baris(n, 1)

}
