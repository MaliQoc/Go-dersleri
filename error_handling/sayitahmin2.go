package errorhandling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {

	aklimdakiSayi := 50

	if tahmin < 1 || tahmin > 100 {
		return "a", errors.New("1-100 arasındaki bir sayıyı giriniz")
	}

	if tahmin > aklimdakiSayi {
		return "Daha küçük bir sayı giriniz.", nil
	}

	if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayı giriniz.", nil
	}

	return "Bildiniz", nil
}

func Demo2() {
	mesaj, _ := TahminEt(101)
	fmt.Println(mesaj) // 1-100 arası sayı girmezseniz a isimli hatayı döndürür. Onun haricinde if içerisindeki değerleri okur.
} // UNUTMAYIN!!! error kütüphanesini kullanarak kendi hata mesajımızı oluşturduk.