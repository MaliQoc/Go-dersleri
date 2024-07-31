package loops

import "fmt"

func SayiTahmin() {
	aklimdakiSayi := 80
	var tahminEdilenSayi int
	var tahminSayisi int

	fmt.Println("Bir sayı tahmin ediniz: ")
	fmt.Scanln(&tahminEdilenSayi)
	tahminSayisi = tahminSayisi + 1

	for aklimdakiSayi != tahminEdilenSayi {

		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Daha büyük bir sayı giriniz.")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}

		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Daha küçük bir sayı giriniz.")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}
	}
	fmt.Println("BİLDİNİZ! Tebrikler.")

	basariDurumu := ""
	if tahminSayisi > 0 && tahminSayisi <= 3 {
		basariDurumu = "Süper"
	} else if tahminSayisi > 3 && tahminSayisi <= 6 {
		basariDurumu = "Güzel"
	} else {
		basariDurumu = "Fena Değil"
	}

	fmt.Printf("%v. tahmin : %v\n", tahminSayisi, basariDurumu)

}
// Sayı tahmin oyunu Go dilinde bu şekilde yazılabilir.