//Wisnu Rananta Raditya Putra (2311102013) IF-11-06

package main
import "fmt"

// Fungsi rekursif untuk menghitung nilai Fibonacci ke-n
func fibonacci(n_2311102013 int) int {

    if n_2311102013 == 0 {
        return 0
    } else if n_2311102013 == 1 {
        return 1
    }
    return fibonacci(n_2311102013-1) + fibonacci(n_2311102013-2)
}

func main() {
    var n int
    fmt.Print("Masukkan nilai n: ")
    fmt.Scan(&n)

    fmt.Println("Deret Fibonacci hingga suku ke-", n)
	for i := 0; i <= n; i++ {
		fmt.Printf("Suku ke-%d: %d\n", i, fibonacci(i))
	}
}
