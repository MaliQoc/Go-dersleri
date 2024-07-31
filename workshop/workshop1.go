package workshop

import "fmt"

func Demos() {
	var a, b, c int = 1, 3, 2
	var enBuyuk int = a

	if enBuyuk < b {
		enBuyuk = b
	}

	if enBuyuk < c {
		enBuyuk = c
	}

	fmt.Println("En büyük sayı: " + fmt.Sprintf("%d", enBuyuk))
}
// Rastgele girdiğimiz 3 sayıdan en büyüğünü ekrana yazdıran fonksiyonu oluşturduk.