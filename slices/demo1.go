package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3) // Dizilerde yeni eleman ekleme çıkarma vs. işlemleri kolaylaştırmak için slice kavramını kullanırız.
	fmt.Println(isimler)

	isimler[0] = "Ali"
	isimler[1] = "Hasan"
	isimler[2] = "Ozan"
	isimler = append(isimler, "Koç") // append diziye ekleme yapar.
	fmt.Println(isimler)
}