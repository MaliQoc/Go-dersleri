package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki para yetersiz!") // if İngilizce'de eğer anlamına gelir. Şartları dile entegre ederken if yapısını kullanırız.
	}

	if cekilmekIstenen <= hesap {
		fmt.Println("Paranız hazırlanıyor...")
		hesap = hesap - cekilmekIstenen
	}

	fmt.Println("Bitti. Hesaptaki para: " + fmt.Sprintf("%f", hesap))
	// fmt.Sprintf ile format işlemini gerçekleştiririz. Parayı float cinsinden yazdırmak istiyoruz. Bu nedenle %f kullandık.
}