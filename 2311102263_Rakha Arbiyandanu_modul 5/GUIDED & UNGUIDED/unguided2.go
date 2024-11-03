package main
import "fmt"

func cetakBintang(n, barisdicetak int){
	if barisdicetak > n {
		return
	}
	for i := 0; i < barisdicetak; i++{
		fmt.Print("*")
	}
	fmt.Println()
	cetakBintang(n, barisdicetak+1)
}

func main(){
	var n int
	fmt.Print("Masukkan jumlah baris :")
	fmt.Scan(&n)
	cetakBintang(n, 1)
}