package main
import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	ratusan := bilangan / 100
	puluhan := (bilangan / 10) % 10
	satuan := bilangan % 10

	fmt.Println(ratusan, puluhan, satuan)
}