//2311102018_ Aryo Tegar Sukarno
package main

import "fmt"

func faktorial(n int ) int {
	if n == 0 || n == 1{
		return 1
	}
	return n * faktorial(n-1)
}
func main(){
	var n int
	fmt.Print("Masukan n mencari Faktorial n!: ")
	fmt.Scanln(&n)
	fmt.Println("Hasil Faktorial dari",n, ":",faktorial(n))
}