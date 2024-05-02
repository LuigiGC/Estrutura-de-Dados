package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var Comida float64
		fmt.Scan(&Comida)

		Dias := DiasRestantes(Comida)
		fmt.Printf("%.0f Dias\n", Dias)
	}
}

func DiasRestantes(Comida float64) float64 {
	Dias := 0.0
	for Comida > 1 {
		Comida /= 2
		Dias++
	}
	return Dias
}
