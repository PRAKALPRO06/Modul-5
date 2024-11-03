package main

import "fmt"

func cetakbintang(win int) {
	if win <= 0 {
		return
	}
	cetakbintang(win - 1)
	for i := 0; i < win; i++ {
		fmt.Print("* ")
	}
	fmt.Println()
}

func main() {
	var win int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&win)

	fmt.Print("")
	cetakbintang(win)
}
