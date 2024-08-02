package maps

import "fmt"

func Demo1() {
	//key - value
	sozluk := make(map[string]string) // map[string]string yaparak, key ve value yu yani anahtar ve değerin hangi değişken tipiyle yazılacağını bildiririz.

	sozluk["Table"] = "Masa"
	sozluk["Book"] = "Kitap"
	sozluk["Pencil"] = "Kalem"

	fmt.Println(sozluk["Book"]) // Kitap sonucunu verir.
	fmt.Println("Eleman sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk) // Sözlüğe kaydedilen eleman sayısını verir. Birnevi dizi yapısıdır.
	delete(sozluk, "Book")          // Book(Kitap) sözcüğünü sildik.
	fmt.Println("Eleman sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)

	deger, varMi := sozluk["Pencil"] // Pencil(Kalem) sözcüğünün olup olmadığını arar. Büyük-küçük harf kullanımı önemlidir. True/false sonucunu etkiler.
	fmt.Println(deger)
	fmt.Println("Listede olma durumu: ", varMi)

	sozluk2 := map[string]string{"Glass": "Bardak", "Microphone": "Mikrofon"}
	fmt.Println(sozluk2)
} // map yapısı C# veya Java gibi dillerdeki dictionary,hash gibi kavramlardan gelir.
// Çoklu dil desteği olan uygulamalarda kullanılır. Sözlük yapısının yerine geçer.