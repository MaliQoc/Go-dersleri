package structs

import "fmt"

type customer struct { // Nesneyi tanımlama sırasında değer atamasını yapıyoruz.
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("Eklendi: ", c.firstName)
}

func (c customer) update() {
	fmt.Println("Güncellendi: ", c.firstName)
}

func Demo2() {
	c := customer{firstName: "Engin", lastName: "Demiroğ", age: 35}
	/* Nesneye özel değişkenleri tanımlarken değişken ismini belirterek de tanımlama yapabiliriz.
	Değişken ismini belirterek atama yaptığımız için sıralamaya dikkat etmemiz gerekli değildir. */
	c.save()
	c.update()
}

/*
Yukarıdaki örnekte nesneyi tanımlama sırasında değer atamasını yaptık.
Nesnenin alt değişkenlerine ulaşarak da tanımlama yapabilirdik.

c := customer{"Engin", "Demiroğ", 35}
c.firstName = "Ali"
c.lastName = "Koç"
c.age = 22
fmt.Println(c) -> Kod {Ali Koç 22} sonucunu verir.
*/