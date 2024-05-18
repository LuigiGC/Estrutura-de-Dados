package main

import (
	"fmt"
	"math"
)

func main() {
	var N, H, C, L int

	for {
		_, err := fmt.Scan(&N)
		if err != nil {
			break
		}
		_, err = fmt.Scan(&H, &C, &L)
		if err != nil {
			break
		}
		alturaTotal := float64(N * H)
		comprimentoTotal := float64(N * C)
		hipotenusa := math.Sqrt(alturaTotal*alturaTotal + comprimentoTotal*comprimentoTotal)


		areaDaRampa := (hipotenusa * float64(L)) / 10000.0

		fmt.Printf("%.4f\n", areaDaRampa)
	}
}

