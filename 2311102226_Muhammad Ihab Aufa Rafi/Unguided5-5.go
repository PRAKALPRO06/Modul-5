package main

import (
	"fmt"
)

func cetakGanjil_226(n int, bilangan int) {
    if bilangan > n {
        return
    }
    
    if bilangan%2 != 0 {
        fmt.Printf("%d ", bilangan)
    }
    cetakGanjil_226(n, bilangan+1)
}

func main() {
    var n int
    
    for {
        fmt.Print("Masukkan bilangan positif (N): ")
        fmt.Scan(&n)
        
        if n > 0 {
            break
        }
        fmt.Println("Masukkan harus berupa bilangan positif!")
    }
    
    fmt.Printf("Bilangan ganjil dari 1 hingga %d: ", n)
    cetakGanjil_226(n, 1)
}