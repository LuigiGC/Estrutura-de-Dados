package main

import "fmt"

func calcularMedia(val[] float64)float64 {
	soma := 0.0
	for _ , num := range val{
		soma += num
	}
	return soma / float64(len(val))
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


func converteCelsiusParaFahrenheit(x float64) float64 {
	return x*1.8 + 32
}
func main() {
	y := 8
	z := 3
	x := 32.0
	p := 90
	val := []float64{10.0,20.0,10.0,20.0}

	fmt.Println(converteCelsiusParaFahrenheit(x), "°F")
	fmt.Println(minhaPotencia(y, z))
	fmt.Println(verificarParidade(p))
	fmt.Println(calcularMedia(val))
}
