package defer_statement

/*

Defer nedir?

Bir fonksiyonun en sonunda çalışan ve çalıştığını garanti ettiğimiz yapılardır.
Fonksiyonun çalışması bittiğinde çalışmasını istediğimiz başka bir fonksiyon varsa defer ile çağırırız.
Önceki fonksiyon yarım kalsa bile defer ile ifade edilmiş fonksiyon tamamlanır.

*/

import "fmt"

func A() {
	fmt.Println("A fonksiyonu çalıştı.")
}

func B() {
	// defer A() fonksiyonda istediğiniz yere koyun. Sonuç değişmeyecek.
	fmt.Println("B fonksiyonu çalıştı.")
	defer A()
	defer C()
	defer D() // Kutuya ilk giren son çıkar mantığı vardır. İlk belleğe alınan en son yazdırılır.
}

func C() {
	fmt.Println("C fonksiyonu çalıştı.")
}

func D() {
	fmt.Println("D fonksiyonu çalıştı.")
}