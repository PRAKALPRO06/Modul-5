package main
import "fmt"

func cetakbintang(n int) {
    if n <= 0 {
        return 
    }
    cetakbintang(n - 1) 
    for i := 0; i < n; i++ {
        fmt.Print("* ") 
    }
    fmt.Println() 
}

func main() {
    var n int
    fmt.Print("Masukkan bilangan n: ")
    fmt.Scan(&n) 

    fmt.Print ("")
	cetakbintang(n) 
}