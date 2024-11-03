package main

import "fmt"

func power(x, y int) int {
    if y == 0 {
        return 1
    } else if y == 1 {
        return x
    } else if y%2 == 0 {
        // Jika pangkat genap, pangkatkan setengah dari pangkat lalu kuadratkan
        temp := power(x, y/2)
        return temp * temp
    } else {
        // Jika pangkat ganjil, kalikan x dengan hasil pangkat (y-1)
        return x * power(x, y-1)
    }
}

func main() {
    var base, exponent int
    fmt.Print("Masukkan bilangan dasar: ")
    fmt.Scanln(&base)
    fmt.Print("Masukkan pangkat: ")
    fmt.Scanln(&exponent)

    result := power(base, exponent)
    fmt.Printf("%d pangkat %d adalah %d\n", base, exponent, result)
}