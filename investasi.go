package main

import (
	"fmt"
)

type Investasi struct {
	hargaBeli, hargaJual, jumlahSaham, totalBeli, totalJual, keuntunganKotor, biayaTransaksi float64
	pajakKeuntungan       float64
	keuntunganBersih      float64
}

func main() {
	var inv Investasi

	fmt.Print("Masukkan harga beli per saham: Rp ")
	fmt.Scan(&inv.hargaBeli)

	fmt.Print("Masukkan harga jual per saham: Rp ")
	fmt.Scan(&inv.hargaJual)

	fmt.Print("Masukkan jumlah saham: ")
	fmt.Scan(&inv.jumlahSaham)

	inv.totalBeli = inv.hargaBeli * inv.jumlahSaham
	inv.totalJual = inv.hargaJual * inv.jumlahSaham
	inv.keuntunganKotor = inv.totalJual - inv.totalBeli

	inv.biayaTransaksi = 0.002 * inv.totalJual

	if inv.keuntunganKotor > 0 {
		inv.pajakKeuntungan = 0.10 * inv.keuntunganKotor
	} else {
		inv.pajakKeuntungan = 0.0
	}

	inv.keuntunganBersih = inv.keuntunganKotor - inv.biayaTransaksi - inv.pajakKeuntungan

	fmt.Println("\nInformasi Investasi Saham:")
	fmt.Printf("Harga Beli: Rp %.2f\n", inv.hargaBeli)
	fmt.Printf("Harga Jual: Rp %.2f\n", inv.hargaJual)
	fmt.Printf("Jumlah Saham: %.0f\n", inv.jumlahSaham)
	fmt.Printf("Total Investasi Awal: Rp %.2f\n", inv.totalBeli)
	fmt.Printf("Total Penjualan: Rp %.2f\n", inv.totalJual)
	fmt.Printf("Keuntungan Kotor: Rp %.2f\n", inv.keuntunganKotor)
	fmt.Printf("Biaya Transaksi (0.2%%): Rp %.2f\n", inv.biayaTransaksi)
	fmt.Printf("Pajak Keuntungan (10%%): Rp %.2f\n", inv.pajakKeuntungan)
	fmt.Printf("Keuntungan Bersih: Rp %.2f\n", inv.keuntunganBersih)
}