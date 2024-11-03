package main

import "fmt"

func fibonacci_226(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci_226(n-1) + fibonacci_226(n-2)
}

func main() {
    fmt.Printf("Deret Fibonacci: \n\n")
    fmt.Print("n  : ")
    for i := 0; i <= 10; i++ {
        fmt.Printf("%4d", i)
    }
    fmt.Print("\nSn : ")
    for i := 0; i <= 10; i++ {
        fmt.Printf("%4d", fibonacci_226(i))
    }
}