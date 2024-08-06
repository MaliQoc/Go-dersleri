package channels

/*
Kanallar (Channels)
Kanallar, Go dilinde asenkron programlama yaparken değer aktarımı yapabileceğimiz hatlardır.
Kanala değer atanması iş parçacığı tarafından bekleneceği için asenkron işlemler arasındaki senkronizasyonu ayarlayabiliriz. Kanallar make() fonksiyonu ile oluşturulur.
Yani bir nevi atama işlemi yapıyoruz. Atama işleminden farkı, kanala atama işlemi yapılana kadar iş parçacığının devam etmemesidir.

*/

import (
	"fmt"
	"time"
)

func CiftSayilar(ciftSayiCn chan int) {

	top := 0

	for i := 0; i <= 10; i += 2 {
		top = top + i
		fmt.Println("Çift sayı toplama çalışıyor")
		time.Sleep(1 * time.Second)
	}
	ciftSayiCn <- top // Kanala değer göndermek için <- işaretini kullanırız.
}

func TekSayilar(tekSayiCn chan int) {

	top := 0

	for i := 1; i <= 10; i += 2 {
		top = top + i
		fmt.Println("Çarpım çalışıyor")
		time.Sleep(1 * time.Second)
	}
	tekSayiCn <- top
}