package main
import "fmt"

func main () {
	var totalBelanja, diskon, totalakhir int
	fmt.Scan(&totalBelanja, &diskon)
	totalakhir = totalBelanja - (totalBelanja * diskon / 100)
	fmt.Println(totalakhir)
}
