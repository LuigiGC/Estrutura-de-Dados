package main

import (
	"fmt"
)

func extrairDiamantes(entrada string) int {
	var pilha []rune
	diamantes := 0
	for _, caractere := range entrada {
		if caractere == '<' {
			pilha = append(pilha, caractere)
		} else if caractere == '>' && len(pilha) > 0 {
			pilha = pilha[:len(pilha)-1]
			diamantes++
		}
	}
	return diamantes
}
func main() {
	var numCasos int
	fmt.Scanln(&numCasos)
	for i := 0; i < numCasos; i++ {
		var entrada string
		fmt.Scanln(&entrada)
		fmt.Println(extrairDiamantes(entrada))
	}
}
