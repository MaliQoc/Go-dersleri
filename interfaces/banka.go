package interfaces

// Bu kodun amacı farklı kredi çeşitleri için ayrı ayrı hesap türleri kullanabilmektir. Banka uygulamalarında interfaceler sıkça kullanılır.

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
} // Bunlar kredi çeşitleri

type CreditCalculator interface {
	Calculate() float64
} // Burası interfaceteki metodun imzası

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
} // Hesaplama formülleri

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal += credits[i].Calculate()
	}

	return paymentTotal
}

/* Demo2() fonksiyonu içerisindeki farklı kredi türlerinin değerlerini CalculateMonthlyPayment(credits []CreditCalculator) içerisinde aldık.
Bunların döngüyle toplamasını yaptık. Değerini döndürdük. */

func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, address: "Ankara"}
	credit2 := Mortgage{rate: 12, creditPaymentTotal: 50000, address: "İstanbul"}
	credit3 := Car{rate: 15, creditPaymentTotal: 60000, carInfo: "Polo"}

	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)

	fmt.Println("Aylık toplam ödeme: ", total) // Ekrana yazdırma işlemini Demo2() fonksiyonunda yaptık.
}

// Görüldüğü gibi şekil hesaplatma fonksiyonuyla aynı yapıy sahip.