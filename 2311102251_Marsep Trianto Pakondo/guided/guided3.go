package main

import "fmt"

func pangkatDua(n int) int  {
	if n == 0 {
		return 1
	}
	return 2 * pangkatDua(n-1)
}

func main()  {
	var n int
	fmt.Print("Masukkan nilai n untuk mencari 2 pangat n: ")
	fmt.Scanln(&n)
	fmt.Println("Hasil 2 Pangkat ", n, " : ", pangkatDua(n))
}