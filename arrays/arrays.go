package arrays

import "fmt"

func Demo1() {

	var sayilar [5]int
	sayilar[2] = 50
	fmt.Println(sayilar) // Direkt bu şekil çalıştırırsak [0,0,0,0,0] sonucunu verir. Çünkü dizinin elemanları default(varsayılan) değer olan 0'ı alır.
	fmt.Println(sayilar[2])

}

func Demo2() {

	var sehirler [5]string
	sehirler[0] = "İstanbul"
	sehirler[1] = "Ankara"
	sehirler[2] = "İzmir"
	sehirler[3] = "Bursa"
	sehirler[4] = "Antalya"

	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	} // Dizide bulunan şehirleri listeler.
}

func Demo3() {
	sayidizisi := [5]int{20, 70, 60, 10, 35}
	fmt.Println(sayidizisi)

	enBuyuk := sayidizisi[0]

	for i := 0; i < len(sayidizisi); i++ {

		if sayidizisi[i] > enBuyuk {
			enBuyuk = sayidizisi[i]
		}
	}
	fmt.Println("En büyük sayı: ", enBuyuk)
} // Dizideki en büyük elemanı ekrana yazdırır.