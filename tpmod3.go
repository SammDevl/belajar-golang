package main
import "fmt"

func main() {
    var A int
    
    fmt.Scan(&A)
    
	ribuan := (A / 1000) % 10
    ratusan := (A / 100) % 10
    puluhan := (A / 10) % 10
    satuan := A % 10

    B := ribuan*1000 + satuan*100 + puluhan*10 + ratusan
    
    
    sum := A + B
    
    fmt.Printf("%d %d\n", B, sum)
}

func alphabet() {
    var cek int32
    fmt.Scanf("%c", &cek)

    fmt.Println((cek >= 'A' && cek <= 'Z') || (cek >= 'a' && cek <= 'z'))
}