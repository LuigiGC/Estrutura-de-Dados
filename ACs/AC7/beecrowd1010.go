package main

import "fmt"

func main() {
	var codP1, qntP1, codP2, qntP2 int
	var valP1, valP2 float64

	fmt.Scanln(&codP1, &qntP1, &valP1)

	fmt.Scan(&codP2, &qntP2, &valP2)

	total := (float64(qntP1) * valP1) + (float64(qntP2) * valP2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
