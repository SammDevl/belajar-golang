package main

import "fmt"

func main() {
	var nama string
	var jumlah int
	var harga, total float64

	fmt.Print("Nama barang: ")
	fmt.Scan(&nama)
	fmt.Print("Jumlah: ")
	fmt.Scan(&jumlah)
	fmt.Print("Harga satuan: Rp ")
	fmt.Scan(&harga)

	total = float64(jumlah) * harga

	fmt.Printf("\nNama: %s\nJumlah: %d\nHarga: Rp %.2f\nTotal: Rp %.2f\n",
		nama, jumlah, harga, total)
}