package main

import "fmt"

//Fazer
func calcularMedia (a,b float64){

}

func verificarParidade (p int){

}
func minhaPotencia (y,z int)int{
	return y * y
}
//Pronto
func converteCelsiusParaFahrenheit(x float64) float64 {
return x*1.8 + 32
}
func main() {
	y := 3
	z :=2
	x := 32.0

	fmt.Println(converteCelsiusParaFahrenheit(x), "°F")
	fmt.Println(minhaPotencia(y,z))
}
