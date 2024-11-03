package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai n untuk mencari 2 pangkat n: ")
	fmt.Scan(&n)
	fmt.Println("Hasil pangkat", n, ":", pangkat(n))

}
func pangkat(n int) int {
	if n == 0 {
		return 1
	} else {
		return 2 * pangkat(n-1)
	}
}
