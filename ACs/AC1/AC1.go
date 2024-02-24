package main

import "fmt"

//Fazer
func calcularMedia(a, b float64) {

}

func verificarParidade(p int) string {
	if p%2 == 0 {
		return "Par"
	} else {
		return "Impar"
	}
}
func minhaPotencia(y, z int) int {
	res := 1
	for i := 0; i < z; i++ {
		res *= y
	}
	return res
}

//Pronto
func converteCelsiusParaFahrenheit(x float64) float64 {
	return x*1.8 + 32
}
func main() {
	y := 8
	z := 3
	x := 32.0
	p := 99

	fmt.Println(converteCelsiusParaFahrenheit(x), "°F")
	fmt.Println(minhaPotencia(y, z))
	fmt.Println(verificarParidade(p))
}
