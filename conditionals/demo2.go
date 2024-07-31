package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki para yetersiz!")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Paranız hazırlanıyor...")
		fmt.Println("Dikkat! Paranızın tamamını çektiniz.") // Başka bir if kullanmamız gerektiğinde else if kullanabiliriz.
	} else {
		fmt.Println("Paranız hazırlanıyor...") // else yapısını if in gerçekleşmediği durum veya durumlarda (İngilizce'de değilse demek) kullanırız.
	}
}