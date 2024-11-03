package main

import "fmt"

// Fungsi rekursif untuk menghitung deret Fibonacci
func fibonacci(n int) int {
        if n <= 1 {
                return n
        }
        return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
        var n int
        fmt.Print("Masukkan nilai n untuk mencari suku ke-n deret Fibonacci: ")
        fmt.Scanln(&n)

        result := fibonacci(n)
        fmt.Println("Suku ke-", n, "deret Fibonacci adalah:", result)
}