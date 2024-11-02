package main 

import "fmt"

func penjumlahan(n int) int {
	if n == 1 {
		return 1
	}
	return 1 + penjumlahan(n-1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n untuk mencari penjumlahan 1 sebanyak n kali: ")
	fmt.Scan(&n)
	fmt.Println("hasil dari penjumlahan 1 sebanyak", n, "kali adalah", penjumlahan(n))
}
