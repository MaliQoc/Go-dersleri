package errorhandling

import "fmt"

type borderException struct {
	parameter int
	message   string
}

func (b *borderException) Error() string {
	return fmt.Sprintf("%d - %s", b.parameter, b.message)
}

func TahminEt2(tahmin int) (string, error) {

	if tahmin < 1 || tahmin > 100 {
		return "", &borderException{tahmin, "Sınırların dışında kaldı."}
	}

	return "Bildiniz", nil
}

/*
Bu kod, TahminEt2 isimli bir fonksiyonla kullanıcının tahmin ettiği sayının belirli sınırlar içinde olup olmadığını kontrol eder.
Eğer tahmin edilen sayı bu sınırlar içinde değilse, bir hata (error) döndürür.
Kodun amacı, bir sayı tahmini oyunu gibi bir senaryoda geçersiz girişleri ele alarak kullanıcıyı bilgilendirmektir.

borderException Struct şu işe yarar: parameter ve message adlı iki alanı olan bir struct tanımlar. Bu struct, özel hata türü olarak kullanılır.
Error fonksiyonu, bu struct'ın bir metodu olarak tanımlanır ve borderException struct'ı bir hata mesajı döndürmek üzere yapılandırılmıştır.
Hata mesajı, parametre değeri ve mesajı birleştirir. */