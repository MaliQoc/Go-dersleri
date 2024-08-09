package stringfunctions

import (
	"fmt"
	s "strings" // strings paketini (s olarak kısaltılmış) kullanırız.
)

func Demo1() {

	isim := "Engin"
	fmt.Println(s.Count(isim, "g"))    // Engin kelimesinde kaç tane g harfi olduğunu sayar ve sonucu ekrana yazdırır.
	fmt.Println(s.Contains(isim, "a")) // Engin kelimesinde a harfinin olup olmadığını kontrol eder.
	fmt.Println(s.Index(isim, "n"))    // Engin kelimesinde n harfinin ilk görüldüğü yerin indeksini döner.
	sonuc := s.Index(isim, "a")        // a harfinin Engin kelimesinde olup olmadığını kontrol eder. Eğer varsa indeksini döner, yoksa -1 döner.

	if sonuc != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}

	fmt.Println(s.ToLower(isim)) // Engin kelimesini tamamen küçük harfe çevirir.
	fmt.Println(s.ToUpper(isim)) // Engin kelimesini tamamen büyük harfe çevirir.

}