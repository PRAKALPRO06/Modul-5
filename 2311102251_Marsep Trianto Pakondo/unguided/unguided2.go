package main

import "fmt"

func printStars(n int) {
	if n == 0 {
		return
	}
	printStars(n-1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Print("\n")
	
}

func main() {
    var N int
    fmt.Print("Masukkan jumlah baris: ")
    fmt.Scan(&N)
    printStars(N)
}
