package main

import (
	"golesson/arrays"
	"golesson/conditionals"
	"golesson/loops"
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
}