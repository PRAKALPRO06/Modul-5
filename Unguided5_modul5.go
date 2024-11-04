package main

import (
	"fmt"
)

func cetakBilganjil(current, x int) {
	if current > x {
		return
	}

	if current%2 != 0 {
		fmt.Print(current, " ")
	}

	cetakBilganjil(current+1, x)
}

func cetakUrutanganjil(x int) {
	if x <= 0 {
		fmt.Println("Masukan harus positif")
		return
	}
	cetakBilganjil(1, x)
	fmt.Println()
}

func main() {
	var x int

	for {
		fmt.Print("Masukkan angka : ")
		_, err := fmt.Scan(&x)

		if err != nil {
			fmt.Println("angka gak pas,masukkan lagi yang lain.")
			var dump string
			fmt.Scanln(&dump)
			continue
		}

		if x == 0 {
			fmt.Println("Program selesai!")
			break
		}

		fmt.Printf("Bilangan ganjil dari 1 sampai %d:\n", x)
		cetakUrutanganjil(x)
	}
}
