package main 

import "fmt"

type Pair struct {
	current, prev1 int
}

func fibo(n int) Pair {
	if n == 2 {
		return Pair{1, 1}
	}
	prev1 := fibo(n-1).current
	prev2 := fibo(n-1).prev1
	current := prev1 + prev2
	return Pair{current, prev1}
}

func main() {
	var n int
	fmt.Print("n: ")
	fmt.Scan(&n)
	current := fibo(n).current
	fmt.Print("Suku ke-", n, " dari barisan fibonacci adalah ", current)
}