package main
import "fmt"

func DeretFibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return DeretFibonacci(n-1) + DeretFibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Cari deret Fibonacci, Masukkan nilai n : ")
	fmt.Scanln(&n)
	fmt.Println("Fibonacci ke", n, ":", DeretFibonacci(n))
}