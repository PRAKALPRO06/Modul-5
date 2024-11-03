// 2311102266_Hanif reyhan zharan adrytona

package main

import (
	"fmt"
)

func ganjil(n, current int) {
	if current <= n {
		fmt.Print(current, " ")
		ganjil(n, current+2)
	}
}
func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)
	ganjil(n, 1)
	fmt.Println()
}
