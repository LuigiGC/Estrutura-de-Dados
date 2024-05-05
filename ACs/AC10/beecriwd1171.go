package main

import (
	"fmt"
	"sort"
)

func cFrequencia(n int, numeros []int) {
	frequencia := make(map[int]int)

	for _, n := range numeros {
		frequencia[n]++

	}
	var sN []int
	for n := range frequencia{
		sN = append(sN, n)
	}
	sort.Ints(sN)
	for _, n := range sN{
		fmt.Printf("%d aparece %d vez(es)\n", n, frequencia[n])
	}
}


func main() {
	var N int
	fmt.Scanln(&N)

	numeros := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&numeros[i])
	}
	cFrequencia(N, numeros)
}
