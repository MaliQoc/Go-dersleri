package loops

import "fmt"

func Demo1() {

	var metin string = "Merhabalar"
	i := 1
	// for yapısı döngülerde kullanılır. Tek tek alt alta uzun uzun verileri yazmaktan bu şekilde kurtuluruz.
	for i <= 5 {
		fmt.Println(metin)
		i++
	}

	fmt.Println("Bitti")
}