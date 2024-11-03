package main

import "fmt"

func baris(bilangan int) {
if bilangan == 1 {
fmt.Print(1)
}else{
fmt.Print(bilangan)
baris(bilangan - 1)
	}
}


func main() {
	var n int
	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&n)
	fmt.Print("Bilangan n ke 1 : ")
	baris(n)
}