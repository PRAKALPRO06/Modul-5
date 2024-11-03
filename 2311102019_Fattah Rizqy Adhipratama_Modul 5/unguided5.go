package main

import "fmt"

func printOddNumbers(n int) {
        if n < 1 {
                return
        }
        printOddNumbers(n - 2)
        fmt.Print(n, " ")
}

func main() {
        var num int
        fmt.Print("Masukkan bilangan: ")
        fmt.Scanln(&num)
        fmt.Println("Barisan bilangan ganjil:")
        printOddNumbers(num)
}