package main

import (
	"fmt"
	"golesson/arrays"
	"golesson/conditionals"
	"golesson/functions"
	"golesson/loops"
	"golesson/maps"
	"golesson/pointers"
	myrange "golesson/range"
	"golesson/slices"
	"golesson/structs"
	"golesson/variables" // main paketinde golesson modülü altında variables paketini çağrıyoruz.
	"golesson/workshop"
)

func main() {
	variables.Demo1() // paket.fonksiyon şeklinde çalıştırılır.
	conditionals.Demo1()
	conditionals.Demo2()
	workshop.Demos()
	loops.Demo1()
	loops.Demo2()
	loops.SayiTahmin()
	loops.AsalSayi()
	loops.ArkadasSayi()
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
}