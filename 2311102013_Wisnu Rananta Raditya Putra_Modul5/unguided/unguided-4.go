//Wisnu Rananta Raditya Putra (2311102013) IF-11-06

package main
import ("fmt")

func printSequence(n, current_2311102013 int) {
	fmt.Printf("%d ", current_2311102013)

	if current_2311102013 == 1 {
		return
	}

	printSequence(n, current_2311102013-1)
	fmt.Printf("%d ", current_2311102013)
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	printSequence(n, n)
	fmt.Println()
}
