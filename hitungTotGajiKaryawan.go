package main

import "fmt"

type Karyawan struct {
	nama       string
	gajiPokok  float64
	tunjangan  float64
	potongan   float64
	totalGaji  float64
}

func main() {
	var k Karyawan

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&k.nama)

	fmt.Print("Masukkan gaji pokok: Rp ")
	fmt.Scan(&k.gajiPokok)

	fmt.Print("Masukkan tunjangan: Rp ")
	fmt.Scan(&k.tunjangan)

	fmt.Print("Masukkan potongan: Rp ")
	fmt.Scan(&k.potongan)

	k.totalGaji = k.gajiPokok + k.tunjangan - k.potongan

	fmt.Println("\nInformasi Karyawan:")
	fmt.Println("Nama:", k.nama)
	fmt.Printf("Gaji Pokok: Rp %.2f\n", k.gajiPokok)
	fmt.Printf("Tunjangan: Rp %.2f\n", k.tunjangan)
	fmt.Printf("Potongan: Rp %.2f\n", k.potongan)
	fmt.Printf("Total Gaji: Rp %.2f\n", k.totalGaji)
}