package stringfunctions

import (
	"fmt"
	s "strings"
)

func Demo2() {

	isim := "Abdülkadir"
	fmt.Println(s.HasPrefix(isim, "Abd"))                                 // s.HasPrefix() fonksiyonu, Abdülkadir kelimesinin "Abd" ile başlayıp başlamadığını kontrol eder.
	fmt.Println(s.HasSuffix(isim, "ab"))                                  // s.HasSuffix() fonksiyonu, Abdülkadir kelimesinin "ab" ile bitip bitmediğini kontrol eder.
	fmt.Println(s.Index(isim, "bd"))                                      // s.Index(isim, "bd") fonksiyonu, Abdülkadir kelimesinde "bd" alt string'inin ilk bulunduğu indeksin pozisyonunu döndürür.
	harfler := []string{"a", "b", "d", "ü", "l", "k", "a", "d", "i", "r"} // slice oluşturuldu :)
	sonuc := s.Join(harfler, "*")                                         // s.Join() fonksiyonu bu slice taki tüm string elemanlarını "*" ile birleştirir ve sonuc değişkenine atar.
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "+", 3)) // s.Replace() fonksiyonu burada sonuc string'inde ilk 3 "*" karakterini "+" karakteri ile değiştirir.
	fmt.Println(s.Split(sonuc, "-"))           // s.Split() fonksiyonu burada sonuc string'ini "-" karakterine göre böler ve bir slice (string dilimi) döndürür.
	fmt.Println(s.Repeat(sonuc, 5))            // s.Repeat(sonuc, 5) fonksiyonu burada sonuc string'ini 5 kez tekrarlar ve birleştirir.

}