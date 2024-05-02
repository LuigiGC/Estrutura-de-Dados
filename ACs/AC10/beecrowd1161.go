package main

import "fmt"
func calcularFatorial(n int )int {
	if n == 0{
		return 1
	}
	return n * calcularFatorial(n - 1)
}

func main(){

	var N1, N2 int

	for {
		_, err := fmt.Scan(&N1, &N2)
		if err != nil {
			break
		}

		FN1 := calcularFatorial(N1)
		FN2 := calcularFatorial(N2)

		total := FN1 + FN2
		fmt.Println(total)

	}
}