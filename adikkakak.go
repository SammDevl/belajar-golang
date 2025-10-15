package main
import "fmt"

func main(){
	    var adik, kakak int
	fmt.Scan(&adik, &kakak)

	selisih := adik - kakak
	if selisih < 0 {
		selisih = -selisih
	}

	if adik == kakak || selisih == 3 || selisih == 6 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}