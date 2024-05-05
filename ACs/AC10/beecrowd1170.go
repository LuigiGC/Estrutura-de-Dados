package main

import "fmt"

func DiasRestantes(Comida float64) int {
	Dias := 0
	for Comida > 1 {
		Comida /= 2
		Dias++
	}
	return Dias
}
func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var Comida float64
		fmt.Scan(&Comida)
		fmt.Printf("%d dias\n", DiasRestantes(Comida))
	}
}
