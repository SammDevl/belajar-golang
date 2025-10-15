package main
import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	jumlahX := a*d + c*b
	jumlahY := b * d
	fmt.Println("penjumlahan:", jumlahX, "/", jumlahY)

	kurangX := a*d - c*b
	kurangY := b * d
	fmt.Println("pengurangan:", kurangX, "/", kurangY)

	kaliX := a * c
	kaliY := b * d
	fmt.Println("perkalian:", kaliX, "/", kaliY)

	bagiX := a * d
	bagiY := b * c
	fmt.Println("pembagian:", bagiX, "/", bagiY)
}