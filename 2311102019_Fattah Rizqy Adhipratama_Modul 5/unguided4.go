package main

import "fmt"

func printSequence(n int) {
        if n == 0 {
                return
        }
        fmt.Print(n, " ")
        printSequence(n - 1)
        fmt.Print(n, " ")
}

func main() {
        var num int
        fmt.Print("Masukkan bilangan: ")
        fmt.Scanln(&num)
        fmt.Println("Barisan bilangan:")
        printSequence(num)
}