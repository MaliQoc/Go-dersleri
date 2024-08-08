package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt") //os paketinde Open metodunu kullanırız ve osya açma işlemini yapar.
	// f dosyanın kendisidir. err hatadır.
	// nil diğer dillerdeki null kelimesinin yerine geçer.

	if err != nil { // Hata sıfırdan farklıysa hata vardır.
		if pErr, ok := err.(*os.PathError); ok { // Eğer hata bir *os.PathError türündeyse, dosya yolu hatasını belirtir.
			fmt.Println("Dosya bulunamadı.", pErr.Path)
			return
		} else { // Diğer bir hata türü söz konusuysa, genel bir hata mesajı verir.
			fmt.Println("Dosya bulunamadı.")
			return
		}
	} else {
		fmt.Println(f.Name()) // Aksi durumda dosyanın ismini yazsın diyoruz.
	} // Biz Demo1 fonkiyonunu main içerisinden çağırıyoruz. Bu nedenle demo1.txt dosyasını dış dizinde (örneğin gokursu klasörğne yazdıysanız gokursu klasörüne) kaydeteümeliyiz.
}

/*
if pErr, ok := err.(*os.PathError); ok { -> Bu satır birkaç önemli işlemi içerir:

Type Assertion (Tip Doğrulama): err.(*os.PathError) ifadesi, err değişkeninin *os.PathError türünde bir hata olup olmadığını kontrol eder.
Eğer err gerçekten bir *os.PathError ise, bu hata pErr değişkenine atanır.

İkili Atama: pErr, ok := kısmı, iki değeri atar:

pErr: Eğer type assertion başarılı olursa, err değişkeni *os.PathError türünde pErr değişkenine atanır.
ok: Type assertion başarılı olup olmadığını belirten bir boolean değeri. Eğer err değişkeni gerçekten *os.PathError türündeyse ok değeri true olur, aksi takdirde false olur.

Koşul Kontrolü: if ifadesi, ok değerini kontrol eder. Eğer ok true ise, err gerçekten bir *os.PathError türündedir ve pErr değişkeni bu hatayı içerir.
Bu durumda if bloğu içerisindeki kod çalıştırılır. Eğer ok false ise, bu koşul sağlanmadığı için else veya else if bloğuna geçilir ya da if bloğu atlanır. */
