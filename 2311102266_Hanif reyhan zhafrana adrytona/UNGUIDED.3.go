// 2311102266_Hanif reyhan zharan adrytona

package main

import "fmt"

func faktor(n, bagi int) {
	if bagi > n {
		return
	}
	if n%bagi == 0 {
		fmt.Printf("%d ", bagi)
	}
	faktor(n, bagi+1)
}
func main() {
	var n int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)
	fmt.Printf("Faktor dari %d adalah: ", n)
	faktor(n, 1)
	fmt.Println()
}
