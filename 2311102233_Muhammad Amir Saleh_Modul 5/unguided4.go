package main

import "fmt"

func deret(n, angka int, naik bool) {
	fmt.Print(angka, " ")
	if !naik && angka == 1 {
		deret(n, 2, true)
	} else if !(naik && angka == n) {
		if !naik {
			deret(n, angka-1, false)
		} else {
			deret(n, angka+1, true)
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	deret(n, n, false)
}
