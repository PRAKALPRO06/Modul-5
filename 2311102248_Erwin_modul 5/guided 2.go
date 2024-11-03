package main

import "fmt"

func jumlahRekrusifn(n int) int {
	if n == 1 {
		return 1

	}
	return n + jumlahRekrusifn(n-1)
}
func main() {
	var n int
	fmt.Println(" masukan nilai n untuk penjumlahan 1 hingga n : ")
	fmt.Scanln(&n)
	fmt.Println(" Hasil Penjumlahan : ", jumlahRekrusifn(n))
}
