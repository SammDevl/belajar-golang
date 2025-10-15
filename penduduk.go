package main
import "fmt"

func main() {
 var penduduk, kelahiran, imigrasi, kematian, emigrasi int
	fmt.Scan(&penduduk, &kelahiran, &imigrasi, &kematian, &emigrasi)

	total := penduduk + kelahiran + imigrasi - kematian - emigrasi
	fmt.Println(total)
}