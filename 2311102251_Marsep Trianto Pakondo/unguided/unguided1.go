package main

import "fmt"

// Fungsi rekursif untuk menghitung Fibonacci
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func printFibo(n int)  {
    for i := 0; i <= n; i++ {
        fmt.Print(fibonacci(i), " ")
    }
}

func main() {
    var n int
    fmt.Print("masukkan N: ")
    fmt.Scan(&n)
    printFibo(n)
}
