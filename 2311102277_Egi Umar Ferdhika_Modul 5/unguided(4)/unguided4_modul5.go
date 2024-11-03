package main

import "fmt"

func cetakUrutan(n int, descending bool) {
	if n < 1 {
		return
	}

	fmt.Printf("%d ", n)

	if descending && n == 1 {
		cetakUrutan(2, false)
		return
	}

	if descending {
		cetakUrutan(n-1, true)
	} else if n < max {
		cetakUrutan(n+1, false)
	}
}

var max int

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("error")
		return
	}

	max = n
	fmt.Print("Hasil: ")
	cetakUrutan(n, true)
	fmt.Println()
}
