package main

import "fmt"

func pangkat(p, g int) int {
	if g == 0 {
		return 1
	}
	if g == 1 {
		return p
	}
	return p * pangkat(p, g-1)
}

func main() {
	var p, g int

	fmt.Print("Masukkan basis dan pangkat: ")
	_, err := fmt.Scan(&p, &g)

	if err != nil {
		fmt.Println("Input tidak valid!")
		return
	}

	hasil := pangkat(p, g)
	fmt.Printf("Hasil %d pangkat %d adalah: %d\n", p, g, hasil)
}
