package interfaces

import "fmt"

func dogrula(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo3() {
	var sayi1 interface{} = "Engin"
	dogrula(sayi1)

	var sayi2 interface{} = "5"
	dogrula(sayi2)

	var sayi3 interface{} = 5
	dogrula(sayi3)
}

/*
Bu kod, bir interface{} türündeki değerin bir int (tam sayı) olup olmadığını kontrol etmek ve bu duruma göre bir çıktı vermek amacıyla yazılmıştır.
Type assertion (tip doğrulama) kavramını uygulamalı olarak gösteren bir örnektir.
Kodun dogrula fonksiyonu, type assertion kullanarak bir interface{} türündeki bir değerin gerçekten bir int olup olmadığını kontrol eder.
sayi, ok := i.(int) bu kısımda uygulanmıştır. Python, C#, Java gibi dillerdeki try-catch yapısına benzer. */