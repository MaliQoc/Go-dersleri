package interfaces

/*

Interface nedir?

Interface, farklı tasarlanmış ama sonuç olarak aynı amaç için kullanılan structlar için bir standart tanımlamasıdır.

Go dilindeki interface için konuşalım.
En güzel özelliği bu struct yapılara metod bağlayabiliyoruz ve birden fazla struct type verisini interface tipleri üzerinden ayırt ederek çağırabiliyoruz yada yönetebiliyoruz.
Çok güzel ve olayı dallandırmıyor. Yok abstract yok interface :)

Örneğin aşağıda şekil alan hesaplatma fonksiyonumuz var. Şekil alanı kısmını interfacee olarak kaydediyoruz. İçerisindeki area() metodu imza gibi bir şey oluyor.
Daha sonra bu imzamızı şekil hesaplatırken kullanacağız. Dikdörtgen ve dairenin alanını ayrı ayrı struct olarak girdik ve özelliklerini verdik.
Daha sonra area() imzamızla iki şeklin dönüş değerlerini tanımladık. Demo fonksiyonunda bu dönüş değeri fonksiyonlarını çağırdık.
En sonunda school fonksiyonunda şeklimizin alanını yazdırdık.

*/

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func school(s shape) {
	fmt.Println("Şeklin alanı: ", s.area())
}

func Demo1() {
	a := rectangle{width: 10, height: 6}
	school(a) // r yi a değişkeniyle çağırdık. Çünkü zaten rectangle içerisinde area() metodunu yani interface i barındırıyor.

	b := circle{radius: 3}
	school(b) /* Aynı durum burada da geçerli. school() fonksiyonunun içerisinde a ve b yi koyarak school(s shape) kısmını tamamladık.
	Aslında bir pointer mantığıyla çalışır. Bir veri diğer bir verinin yerini tutar. */
}