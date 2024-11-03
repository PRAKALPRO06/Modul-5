package main

import "fmt"

func faktorial ( n int) int {
	if n==0 || n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}

func main (){
	var n int
	fmt.Print("Masukkan nilai n untuk mencari faktorial n!: ")
	fmt.Scanln(&n)
	fmt.Println("Hasil faktorial dari", n, ":", faktorial(n))
}