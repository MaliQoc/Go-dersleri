package slices

import "fmt"

func Demo2() {
	sehirler := []string{"İstanbul", "Ankara", "İzmir"}
	fmt.Println(sehirler)

	sehirlerKopya := make([]string, len(sehirler))
	fmt.Println(sehirlerKopya)

	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya) // şehirler dizisinin elemanları sehirlerKopya dizisine kopyalanmış oldu.

	sehirler = append(sehirler, "Bursa")
	fmt.Println(sehirler) // sehirler dizisine bursayı da ekledik.
	fmt.Println(sehirlerKopya)

	fmt.Println(sehirler[1:3]) // Dizide 1. indisten 3. indise kadar olan kısmı alacak.
	fmt.Println(sehirler[1:])  // Dizide 1. indisten son indise kadar olan kısmı aldı.
	fmt.Println(sehirler[:2])  // Dizide 2. indise kadar olan kısmı aldı.
}