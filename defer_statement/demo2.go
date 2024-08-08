package defer_statement

import "fmt"

func Demo2(sayi int32) string {

	defer DeferFunc()

	if sayi%2 == 0 {
		fmt.Println("Çift sayı çalıştı.")
		return "Çift sayıdır." // Önce çift sayı çalıştı yazdırır. Sonra defer ü yazdırır. En sonda return değerini görüp döndürür.
	} // DeferFunc fonksiyonunu aynı fonksiyon içerisinde başta yazmamız gerekiyor. Derleyici return ü görünce onu yazdırır.

	if sayi%2 != 0 {
		return "Tek sayıdır."
	}

	return "Belli değil"
}

func Test() {
	sonuc := Demo2(10)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti.")
}
