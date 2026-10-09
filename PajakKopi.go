package main
import "fmt"

func main () {
	const hargaKopi int = 20000
	const pajak float64 = 0.11

	var jumlahKopi int
	fmt.Scan(&jumlahKopi)
	totalHargaTanpaPajak := float64(jumlahKopi * hargaKopi)
	totalPajak := totalHargaTanpaPajak * pajak
	totalPembayaran := int(totalHargaTanpaPajak + totalPajak)
	fmt.Println(totalPembayaran)
}