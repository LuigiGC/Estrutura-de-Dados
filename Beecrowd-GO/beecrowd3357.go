package main

import "fmt"

func main() {

	var N int
	var L, Q float64
	var part []string
	
	fmt.Scanln(&N, &L, &Q)

	part = make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&part[i])
	}
	cont := 0
	for L > Q {
		L = L - Q
		cont += 1
		if cont == N {
			cont = 0
		}

	}
	res := part[cont]

	fmt.Printf("%s %.1f\n", res, L)

}
