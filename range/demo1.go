package myrange // Range anahtar bir kelime olduğu için paket ismimizi değiştirdik. range yerine Range yazarak da bunu yapabilirdik.

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"}

	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}

	for i, sehir := range sehirler {
		fmt.Print(i + 1)
		fmt.Print("- ")
		fmt.Println(sehir)
	} // Yukarıdaki for döngüsüyle for_range döngüsü aynı şeyi yazdırır fakat burada daha az kod vardır ve dizi indis numaraları başa yazdırılır.
}