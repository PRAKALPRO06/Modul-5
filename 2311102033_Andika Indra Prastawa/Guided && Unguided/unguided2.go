package main

import "fmt"

// Fungsi rekursif untuk mencetak bintang di setiap baris
func printStars(n int) {
	if n > 0 {
		printStars(n - 1)
		for i := 0; i < n; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scanln(&n)

	fmt.Println("Pola Bintang:")
	printStars(n)
}
