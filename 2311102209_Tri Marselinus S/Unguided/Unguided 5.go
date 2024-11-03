package main

import "fmt"


func cetakGanjil(n int, a int) {

    if a > n {
        return
    }
    if a%2 != 0 {
        fmt.Print(a, " ")
    }
    cetakGanjil(n, a+1)
}

func main() {
    var n int
    fmt.Print("Masukkan bilangan bulat positif N: ")
    fmt.Scan(&n)

    fmt.Print("")
    cetakGanjil(n, 1) 
}
