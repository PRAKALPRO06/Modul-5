package main

import "fmt"


func cetakBarisan(n int, batas int) {
    fmt.Print(batas, " ") 
    if batas == 1 {
        return
    }
    cetakBarisan(n, batas-1)
    fmt.Print(batas, " ")
}

func main() {
    var n int
    fmt.Print("Masukkan bilangan bulat positif N: ")
    fmt.Scan(&n)

    fmt.Print("Barisan dari ", n, " adalah: ")
    cetakBarisan(n, n) 
}
