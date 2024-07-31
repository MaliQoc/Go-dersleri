package loops

import "fmt"

func ArkadasSayi() {
	a := 213
	b := 324
	top1 := 0
	top2 := 0

	for i := 1; i < a; i++ {
		if a%i == 0 {
			top1 += i
		}
	}

	for i := 1; i < b; i++ {
		if b%i == 0 {
			top2 += i
		}
	}

	if top1 == b && top2 == a {
		fmt.Printf("%v ve %v arkadaş sayılardır.\n", a, b)
	} else {
		fmt.Printf("%v ve %v arkadaş sayı değillerdir.\n", a, b)
	}
}

// Girdiğimiz iki sayının arkadaş sayı olup olmadığını ölçmüş olduk.
