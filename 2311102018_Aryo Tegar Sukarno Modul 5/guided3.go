//2311102018_ Aryo Tegar Sukarno
package main

import "fmt"

func pangkatdua(n int) int {
	if n == 0 {
		return 1 
	}
	return 2 * pangkatdua(n-1)
} 

func main() {
	var n int 
	fmt.Print("Masukkan nilai n untuk mencari 2 pangkat n: ")
	fmt.Scan(&n) // Perbaiki 'fmt.scan' menjadi 'fmt.Scan'
	fmt.Println("Hasil 2 Pangkat", n, ":", pangkatdua(n))
}