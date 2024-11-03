package main

import (
	"fmt"
)

func cariFaktor_226(n int, bilangan int) {
    if bilangan > n {
        return
    }
    
    if n%bilangan == 0 {
        fmt.Printf("%d ", bilangan)
    }
    
    cariFaktor_226(n, bilangan+1)
}

func main() {
    var n int
    
    for {
        fmt.Print("Masukkan bilangan positif (n): ")
        fmt.Scan(&n)
        if n > 0 {
            break
        }
        fmt.Println("Masukkan harus berupa bilangan positif!")
    }
    fmt.Printf("Faktor-faktor dari %d adalah: ", n)
    cariFaktor_226(n, 1)
}