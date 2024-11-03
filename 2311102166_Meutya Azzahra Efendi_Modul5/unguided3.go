// Meutya Azzahra Efendi
// 2311102166
// IF-11-06
package main

import "fmt"

func isFaktor(bilangan, faktor int) bool {
        return bilangan%faktor == 0
}

func cetakFaktor(bilangan, faktor int) {
        if faktor > bilangan {
                return
        }
        if isFaktor(bilangan, faktor) {
                fmt.Print(faktor, " ")
        }
        cetakFaktor(bilangan, faktor+1)
}

func main() {
        var bilangan int
        fmt.Print("Masukkan bilangan: ")
        fmt.Scanln(&bilangan)

        fmt.Printf("Faktor dari %d adalah: ", bilangan)
        cetakFaktor(bilangan, 1)
        fmt.Println()
}