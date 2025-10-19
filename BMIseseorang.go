package main

import "fmt"

type DataBMI struct {
	nama   string
	berat  float64
	tinggi float64
	bmi    float64
}

func main() {
	var data DataBMI

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&data.nama)

	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scan(&data.berat)

	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scan(&data.tinggi)

	data.bmi = data.berat / (data.tinggi * data.tinggi)

	fmt.Println("\nInformasi BMI:")
	fmt.Println("Nama:", data.nama)
	fmt.Printf("Berat: %.1f kg\n", data.berat)
	fmt.Printf("Tinggi: %.2f m\n", data.tinggi)
	fmt.Printf("BMI: %.2f\n", data.bmi)
}