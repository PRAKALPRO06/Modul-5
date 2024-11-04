package main

import (
	"fmt"
)

func cetakTurun(n int) {
	if n == 1 {
		fmt.Print(n, " ")
		return
	}
	fmt.Print(n, " ")
	cetakTurun(n - 1)
}

func cetakNaik(n, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	cetakNaik(n, current+1)
}

func cetakKeseluruhan(n int) {
	if n <= 0 {
		fmt.Println("Masukan harus berupa bilangan bulat positif!")
		return
	}
	cetakTurun(n)
	cetakNaik(n, 2)
	fmt.Println()
}

func main() {
	var n int

	for {
		fmt.Print("Masukkan angka,ketik 0 kalau selsai: ")
		_, err := fmt.Scan(&n)

		if err != nil {
			fmt.Println("angka kurang pas,coba lagi masukkan yang lain.")
			var dump string
			fmt.Scanln(&dump)
			continue
		}

		if n == 0 {
			fmt.Println("Program selesai!")
			break
		}

		fmt.Printf("Barisan bilangan untuk N = %d:\n", n)
		cetakKeseluruhan(n)
	}
}
