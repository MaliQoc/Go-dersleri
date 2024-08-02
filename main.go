package main

import (
	"fmt"
	"golesson/arrays"
	"golesson/conditionals"
	"golesson/functions"
	"golesson/loops"
	"golesson/maps"
	myrange "golesson/range"
	"golesson/slices"
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
}