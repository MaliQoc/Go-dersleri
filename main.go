package main

import (
	"fmt"
	"golesson/arrays"
	"golesson/defer_statement"
	errorhandling "golesson/error_handling"
	stringfunctions "golesson/string_functions"

	// "golesson/channels"
	"golesson/conditionals"
	"golesson/functions"

	// "golesson/goroutines"
	"golesson/interfaces"
	"golesson/loops"
	"golesson/maps"
	"golesson/pointers"
	myrange "golesson/range"
	"golesson/slices"
	"golesson/structs"
	"golesson/variables" // main paketinde golesson modülü altında variables paketini çağrıyoruz.
	"golesson/workshop"
	// "time"
)

func main() {

	variables.Demo1() // paket.fonksiyon şeklinde çalıştırılır.
	conditionals.Demo1()
	conditionals.Demo2()
	workshop.Demos()
	loops.Demo1()
	loops.Demo2()
	// loops.SayiTahmin()
	// loops.AsalSayi()
	loops.ArkadasSayi()
	loops.SwitchDemo1()
	arrays.Demo1()
	arrays.Demo2()
	arrays.Demo3()
	arrays.Multiarrays()
	slices.Demo1()
	slices.Demo2()

	functions.SelamVer("ali")

	var sonuç = functions.Topla(5, 8)
	fmt.Println(sonuç * 10)

	sonuç1, sonuç2, sonuç3, sonuç4 := functions.DortIslem(10, 2)
	fmt.Println("Toplam: ", sonuç1)
	fmt.Println("Çıkarma: ", sonuç2)
	fmt.Println("Çarpım: ", sonuç3)
	fmt.Println("Bölüm: ", sonuç4)

	fmt.Println(functions.ToplaVariadic(1, 4, 6, 3, 10))
	fmt.Println(functions.ToplaVariadic(1, 4))
	fmt.Println(functions.ToplaVariadic())

	sayilar := []int{4, 6, 7, 0, 11}
	fmt.Println(functions.ToplaVariadic(sayilar...)) // Yukarıdaki ve aşağıdaki grup gibi 2 kullanım vardır.

	maps.Demo1()
	myrange.Demo1()
	myrange.Demo2()

	sayi := 20
	pointers.Demo1(sayi)
	fmt.Println("Maindeki sayı: ", sayi) // Burada sayıyı normal bir şekilde yazdıracaktır. Toplanmış hali demoda bulunduğu için orada artar.

	pointers.Demo2(&sayi)
	fmt.Println("Maindeki sayı: ", sayi) // Burada sayının artırılmış halini yazdırır. Main kısmında da yeni değer 21 olmuştur.

	sayilar = []int{1, 2, 3}
	pointers.Demo3(sayilar)
	fmt.Println("Maindeki sayı: ", sayilar[0]) /* Arrayler yani diziler diğer değişkenler gibi değerler üzerinden çalışmazlar.
	Referans dediğimiz bellekteki adreslerle çalışır. Bu nedenle maindeki değerlerde de değişim olacaktır.
	Biz array e verileri adresiyle gönderiyoruz. Slicelarda da aynı durum vardır. */

	structs.Demo1()
	structs.Demo2()

	/*
		goroutines.CiftSayilar()
		goroutines.TekSayilar()

		// Eşzamanlı bir işlem oluşturmak için go anahtar kelimesinden faydalanabiliriz. Bunun için eşzamanlı çalışacak işlemin başına go yazmamız yeterli olacaktır.

		go goroutines.CiftSayilar()
		go goroutines.TekSayilar()
		time.Sleep(3 * time.Second)
		fmt.Println("Main bitti.")

		/* time.Sleep() kullanarak 3 saniye bekletiyoruz.
		Eğer time.Sleep() eklememiş olsaydık ekranda önceki fonksiyonlar çağrılacak, sıra bu fonksiyona geldiğinde ise program bunu atlayacaktı.
		Bunun sebebi Go Runtime'ının sadece ana iş parçacığını beklemesidir. Ana iş parçacığındaki işlemler sonlandıktan sonra, diğer işlemleri beklemiyor.
		Yukarıdaki örnekte bunu engellemek için time.Sleep() kullandık. Böylece program 5 saniye beklerken eşzamanlı işlemimiz de tamamlandı. */

	/*
		ciftSayiCn := make(chan int) // Bu satırda make() fonksiyonu ile ciftSayiCn isminde bir kanal oluşturduk. Bu kanalın özelliği int tipinde değer taşımasıdır.
		tekSayiCn := make(chan int)  // Bu satırda da aynı işlemi yaptık. Kanalımızın ismi tekSayiCn.
		go channels.CiftSayilar(ciftSayiCn)
		go channels.TekSayilar(tekSayiCn)

		ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn /* Bir de bu kanalın çıkış noktası olması gerekir.
		Bu çıkış noktasında, ister kanaldan gelen veriyi bir değişkene atayabiliriz, istersek de sadece kanala veri gelmesini bekleyebiliriz. */

	/* Örnekte ciftSayiToplam isimli değişkene ciftSayiCn kanalından gelen int tipinde değer atadık.
	Bir de tekSayiToplam isimli değişkene tekSayiCn kanalından gelen int tipinde değer atadık.
	Bu iki değişkene atama işlemi iki kanala değer gönderildiği zaman yapılacaktır.
	Yani ciftSayiCn ve tekSayiCn kanallarına değer gelene kadar iş parçacığı duraklatılacaktır. */

	/* carpim := ciftSayiToplam * tekSayiToplam
	fmt.Println("Çarpım: ", carpim) */

	interfaces.Demo1()
	interfaces.Demo2()
	defer_statement.B()
	defer_statement.Test()
	defer_statement.Demo3()
	errorhandling.Demo1()
	interfaces.Demo3()
	errorhandling.Demo2()
	fmt.Println(errorhandling.TahminEt2(102))
	stringfunctions.Demo1()
	stringfunctions.Demo2()
}