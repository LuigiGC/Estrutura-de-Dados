package main

import "fmt"

func honoi(n int, origem, destino, aux string) {
	if n == 1 {
		fmt.Println("Mova o disco de cima de ", origem, " Para ", destino)
		return
	}
	honoi(n-1, origem, aux, destino)
	fmt.Println("Mova o disco ", n, " de ", origem, " Para", destino)
	honoi(n-1, aux, destino, origem)
}

func verificarSoma(listaNumeros []int, x int) (int, int) {
	ini := 0
	fin := len(listaNumeros) - 1

	for ini < fin {
		sum := listaNumeros[ini] + listaNumeros[fin]

		if sum == x {
			return listaNumeros[ini], listaNumeros[fin]
		} else if sum < x {
			ini++
		} else {
			fin--
		}
	}
	return -1, -1
}
func main() {
	discos := 4
	honoi(discos, "Torre 1", "Torre 3", "Torre 2")

	listaNumeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := 11

	num1, num2 := verificarSoma(listaNumeros, x)

	if num1 == -1 && num2 == -1 {
		fmt.Printf("Nenhum par para o valor de %d foi encontrado", x)

	} else {
		fmt.Println("Soma de pares: ", num1, "+", num2, "=", x)
	}

}
