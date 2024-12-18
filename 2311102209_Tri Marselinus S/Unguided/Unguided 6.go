package main

import "fmt"


func pangkat(x int, y int) int {
    if y == 0 {
        return 1
    }
    if y < 0 {
        return 1 / pangkat(x, -y)
    }
    return x * pangkat(x, y-1)
}

func main() {
    var x, y int
    fmt.Print("Masukkan bilangan bulat x: ")
    fmt.Scan(&x)
    fmt.Print("Masukkan bilangan bulat y: ")
    fmt.Scan(&y)

    hasil := pangkat(x, y)
    fmt.Print(hasil)
}
