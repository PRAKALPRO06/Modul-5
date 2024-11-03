package main

import "fmt"

func printStars(rows int) {
	if rows <= 0 {
		return
	}
	printStars(rows - 1)
	for col := 0; col < rows; col++ {
		fmt.Print("* ")
	}
	fmt.Println()
}

func main() {
	var numRows int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&numRows)

	printStars(numRows)
}
