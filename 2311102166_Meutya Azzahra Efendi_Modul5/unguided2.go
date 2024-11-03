// Meutya Azzahra Efendi
// 2311102166
// IF-11-06
package main

import "fmt"

func printStars(n int) {
	if n == 0 {
		return
	}
	printStars(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris bintang: ")
	fmt.Scanln(&n)

	printStars(n)
}
