package arrays

import "fmt"

func Multiarrays() {
	var sayidizisi [2][3]int

	sayidizisi[0][0] = 1
	sayidizisi[0][1] = 3
	sayidizisi[0][2] = 5
	sayidizisi[1][0] = 2
	sayidizisi[1][1] = 4
	sayidizisi[1][2] = 6

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(sayidizisi[i][j])
			fmt.Print(" ")
		}
		fmt.Println() // Matris görünümü için bu şekil bir kullanım yaptık. Sütunda aralara boşluk + satır-sütun ayrımı için alta bir boşluk koyduk.
	}
}