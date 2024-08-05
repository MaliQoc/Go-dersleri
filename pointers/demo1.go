package pointers

/*

Pointer’lar, verilerin kendilerini değil, verilerin nerede bulunduğu bilgisini taşır.
Bu sayede programın bellek kullanımını optimize etmemize ve kodumuzun daha verimli çalışmasını sağlar.
Pointer denildiğinde * ve & operatörleri aklımıza gelir.
& operatörü değişken adreslerini bulmakta kullanılır.
* operatörü ile pointerın işaret ettiği değer alınır.

*/

import (
	"fmt"
)

func Demo1(sayi int) {
	sayi += 1
	fmt.Println("Demodaki sayı: ", sayi)
}

func Demo2(sayi *int) {
	*sayi += 1
	fmt.Println("Demodaki sayı: ", sayi)  // sayi değişkeninin bellekteki adresini verir.
	fmt.Println("Demodaki sayı: ", *sayi) // Bellek adresindeki değeri verir.
}

func Demo3(sayilar []int) {
	sayilar[0] = 100
	fmt.Println("Demodaki sayı: ", sayilar[0])
}