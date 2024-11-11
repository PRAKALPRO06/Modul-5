package main

import "fmt"

func faktor(n, i int) {
	if i <= n {
		if n%i == 0 {
			fmt.Println(i)
		}

		faktor(n, i+1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n untuk mencari faktor-faktornya: ")
	fmt.Scanln(&n)
	fmt.Println("Faktor-faktor dari", n, "adalah:")
	faktor(n, 1)
}
