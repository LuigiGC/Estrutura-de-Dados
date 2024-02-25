package main

import "fmt"

func calcularMedia(val ... float64)float64 {
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
func minhaPotencia(base, expoente int) int {
	res := 1
	for i := 0; i < expoente; i++ {
		res *= base
	}
	return res
}


func converteCelsiusParaFahrenheit(temp float64) float64 {
	return temp*1.8 + 32
}
func main() {
	base := 8
	expoente := 3
	temp := 32.0
	p := 90


	fmt.Println(converteCelsiusParaFahrenheit(temp), "°F")
	fmt.Println(minhaPotencia(base, expoente))
	fmt.Println(verificarParidade(p))
	fmt.Println(calcularMedia(10.0,20.0,10.0,15.0))
}