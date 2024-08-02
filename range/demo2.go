package myrange

import "fmt"

func Demo2() {
	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	top := 0

	for _, sayi := range sayilar {
		if sayi%2 != 0 {
			top = top + sayi
		}
	}
	fmt.Println("Toplam: ", top)
}