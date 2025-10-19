package main

import "fmt"

type PersegiPanjang struct {
	panjang, lebar, luas, keliling float64
}

func main() {
	var p PersegiPanjang

	fmt.Print("Masukkan panjang: ")
	fmt.Scan(&p.panjang)

	fmt.Print("Masukkan lebar: ")
	fmt.Scan(&p.lebar)

	p.luas = p.panjang * p.lebar
	p.keliling = 2 * (p.panjang + p.lebar)

	fmt.Println("\nInformasi Persegi Panjang:")
	fmt.Printf("Panjang: %.1f\n", p.panjang)
	fmt.Printf("Lebar: %.1f\n", p.lebar)
	fmt.Printf("Luas: %.1f\n", p.luas)
	fmt.Printf("Keliling: %.1f\n", p.keliling)
}