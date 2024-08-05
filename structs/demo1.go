package structs

/* Go programlama dilinde C# ve Java'da olduğu gibi sınıflar (class) yoktur. Sınıflar yerine struct'lar (yapılar) vardır.
Structlar sayesinde bir nesne oluşturabilir ve bu nesneye ait özellikler oluşturabilirsiniz.
Farklı veri türlerine sahip verileri bir arada tutabilir ve bu veriler üzerinde işlemler yapabilirsiniz. */

import "fmt"

func Demo1() {
	fmt.Println(product{"Bilgisayar", 500, "XYZ"})
	fmt.Println(product{"Bilgisayar", 500, "ABC"})
	fmt.Println(product{name: "Bilgisayar", unitPrice: 500})
}

// type yeni bir tür oluşturmak içindir.
type product struct {
	name      string
	unitPrice float64
	brand     string
}