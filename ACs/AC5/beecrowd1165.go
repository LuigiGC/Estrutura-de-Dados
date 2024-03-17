package main

import "fmt"

func main() {

	var T, PA, PB, i int
	var G1, G2 float64
	cont := 0

	fmt.Scanln(&T)

	for i = 0; i < T; i++ {
		fmt.Scanln(&PA, &PB, &G1, &G2)

		for PA <= PB && cont <= 100 {
			PA += int(float64(PA) * (G1 / 100))
			PB += int(float64(PB) * (G2 / 100))

			cont++
		}
		if cont > 100 {
			fmt.Println("Mais de 1 seculo.")
		} else {
			fmt.Printf("%d anos.\n", cont)
		}
		cont = 0
	}
}