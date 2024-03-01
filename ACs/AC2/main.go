package main

import (
	"fmt"
	"math"
	geometria "projeto/retangulo"
)

func inverterString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	vetor := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	for _, elemento := range vetor {
		fmt.Println(elemento)
	}

	var entrada string
	fmt.Println("Digite algo para inverter: ")
	fmt.Scanln(&entrada)

	stringInvertida := inverterString(entrada)

	fmt.Println("Palavra invertida: ", stringInvertida)

	ponto := Ponto{X: 3, Y: 4}

	distancia := ponto.DistanciaOrigem()

	fmt.Println("A distância do ponto até a origem é \n", distancia)

	var base, altura float64
	fmt.Println("Digite a base do retângulo:")
	fmt.Scanln(&base)

	fmt.Println("Digite a altura do retângulo:")
	fmt.Scanln(&altura)

	area := geometria.AreaRetangulo(base, altura)
	perimetro := geometria.PerimetroRetangulo(base, altura)

	fmt.Println("Área do retângulo: ", area)
	fmt.Println("Perímetro do retângulo: ", perimetro)

}

type Ponto struct {
	X float64
	Y float64
}

func (p Ponto) DistanciaOrigem() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
