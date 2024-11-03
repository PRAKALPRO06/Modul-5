//2311102018_ Aryo Tegar Sukarno
package main

import "fmt"

func jumlahrekursif(n int) int {
	if n == 1{
		return 1 
	}
	return n + jumlahrekursif(n-1)
}

func main() {
	var n int
	fmt.Print("Masukan nilai n untuk penjumlahan 1 hingga n: ")
	fmt.Scanln(&n)
	fmt.Println("Hasil penjumlahan : " , jumlahrekursif(n)) 
}