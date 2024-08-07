package loops

import "fmt"

/* Switch kelimesinin Türkçe’deki anlamı anahtardır. Switch deyimi de if-else deyimi gibi koşul üzerine çalışır.
Switch’te koşulların gerçekleşmediği zaman işlem uygulamak istiyorsak bunu default terimi ile yaparız. */

func SwitchDemo1() {

	i := 5

	switch i {
	case 5:
		fmt.Println("i eşittir 5")
	case 10:
		fmt.Println("i eşittir 10")
	case 15:
		fmt.Println("i eşittir 15")
	default:
		fmt.Println("i 5, 10 ve 15 değil.")
	}

	// Aşağıdaki yöntemde switch deyiminin yanına koşul girmek yerine case deyiminin yanına koşul giriyoruz.

	/* func main() {
		x := 5

	switch {
	case x == 5:
		fmt.Println("x 5'tir")
		fallthrough // Durumlar içerisinde kontrol etmemiz gereken başka durumlarda olabilir. Bunun için fallthrough deyimini kullanabiliriz.
	case x < 10:
		fmt.Println("x 10'dan küçüktür")
	} */
}