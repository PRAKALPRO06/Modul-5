package main

import "fmt"

func penjumlahan(n int)int{
	if n == 1 {
		return 1
	}else{
		return n + penjumlahan(n-1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&n)
	fmt.Print("Hasil : ")
	fmt.Println(penjumlahan(n))
}