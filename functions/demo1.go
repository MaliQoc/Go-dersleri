package functions

import (
	"fmt"
)

func Topla(a int, b int) int {
	var top = a + b
	return top
} // Topla fonksiyonunu main kısmında returnledik yani çağırdık.

func SelamVer(kullaniciAdi string) string {
	fmt.Println("Merhaba", kullaniciAdi) // main kısmında bu fonksiyonu çağırırken ne yazarsanız yazın başına merhaba ekleyecektir.
	return kullaniciAdi
}

func DortIslem(s1 int, s2 int) (int, int, int, float32) {

	toplama := s1 + s2
	cikarma := s1 - s2
	carpma := s1 * s2
	bolme := float32(s1 / s2) // Bölme işleminde float32 parantezine almak gerekir.

	return toplama, cikarma, carpma, bolme // Bu kısımda çoklu fonksiyonları gördük.
}

func ToplaVariadic(sayilar ...int) int {
	top := 0
	for i := 0; i < len(sayilar); i++ {
		top = top + sayilar[i]
	}
	return top // Diziye istediğimiz kadar elemanı sığdırmak için böyle bir kullanıma daha sahibiz. Bu kullanıma variadic fonksiyon diyoruz.
}