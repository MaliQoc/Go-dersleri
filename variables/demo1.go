package variables

import "fmt" // Kullanılmayan kütüphaneler hata verir.

func Demo1() {
	fmt.Print("Merhaba")   // Print gibi komutlar büyük harfle başlamalıdır. Go dili büyük küçük harf duyarlıdır.
	fmt.Println(" Dünya!") // Print yazdırır, Println alt satıra yazdırır.

	var metin string = "Merhaba Türkiye!"
	fmt.Println(metin)

	var kdv int = 10
	fmt.Println(kdv)
	fmt.Println(100 * 10 / 200)

	var kdv2 float32 = 0.1
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)

	// Yukarıda hep var ile değişken ürettik. Yani tipini biz belirledik.

	kdv3 := 25 // Burada ise verdiğimiz değerin hangi tür olduğunu Go dili kendisi yazdırıyor.
	fmt.Println(kdv3)
	fmt.Printf("Veri tipi: %T\n", kdv3) // Formatlı yazdırmayı kullanıyoruz. Bu nedenle Printf kullandık.

	var durum bool
	var durum2 bool

	var metin1 string = "Engin"
	var metin2 string = "Ahmet"

	durum = metin1 == metin2
	durum2 = metin1 != metin2

	fmt.Println(durum)
	fmt.Println(!durum)
	fmt.Println(durum2) // durum2 değişkeniyle hem false sonucunu verecek durum değişkeninin tersini yazmış olduk hem de bunu yazdırırken ünlem işaretiyle yaptık.

}

/* := yeni bir değişken bildirimi ve ataması için kullanılır.
= mevcut bir değişkene yeni bir değer atamak için kullanılır. */
