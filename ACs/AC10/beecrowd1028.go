package main

import "fmt"

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func main() {
	var Teste, N1, N2 int

	fmt.Scan(&Teste)

	for i := 0; i < Teste; i++ {
		fmt.Scan(&N1, &N2)
		tamanhoMaxPilha := mdc(N1, N2)
		fmt.Println(tamanhoMaxPilha)
	}
}
