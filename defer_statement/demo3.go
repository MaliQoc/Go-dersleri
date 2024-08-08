package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Ürün kaydedildi: ", p.productName)
	defer Log() // Log fonksiyonunun her zaman çalıştığını görmek için defer koyduk. Hata yakalamalarda log çalışırken hatanın neden olduğunu görebiliriz.
}

func Log() {
	fmt.Println("Loglandı.")
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 30000}
	defer p.save()
	p = product{productName: "Mouse", unitPrice: 1000}
	fmt.Println("İşlem başarılı.") // Bu senaryoda laptopu defer ifadeden sonra kaydeder.
} // deferi mouse un altına ekleseydik mouse u kaydederdi. deferin üstünde ne varsa o okunur ve kaydolur sonraki atamalar işe yaramaz.