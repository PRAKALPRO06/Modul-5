package main

import (
	"fmt"
)

func cetakTurun_226(n int, bilangan int) {
    if bilangan < 1 {
        return
    }
    fmt.Printf("%d ", bilangan)
    cetakTurun_226(n, bilangan-1)
}

func cetakNaik_226(n int, bilangan int) {
    if bilangan > n {
        return
    }
    fmt.Printf("%d ", bilangan)
    cetakNaik_226(n, bilangan+1)
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
    
    fmt.Printf("Barisan bilangan untuk N = %d: ", n)
    cetakTurun_226(n, n)
    cetakNaik_226(n, 2)
}