package main

import "fmt"

func tomadas() {
	var t1, t2, t3, t4 int8
	fmt.Scanln(&t1, &t2, &t3, &t4)
	fmt.Println(t1, t2, t3, t4)

}

func meteoritos() {
	var x1, x2, y1, y2, x, y, n int
	num := make([]int, 0)
	i := 0
	for {
		fmt.Scanln(&x1, &y1, &x2, &y2)
		if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 {
			break
		}
		fmt.Scanln(&n)
		num = append(num, 0)
		for j := 1; j <= n; j++ {
			fmt.Scanln(&x, &y)
			if x1 <= x && x <= x2 && y2 <= y && y <= y1 {
				num[i]++
			}
		}
		i++
	}
	for j, num_meteoritos := range num {
		fmt.Println("Teste", j+1)
		fmt.Println(num_meteoritos)
		fmt.Println("")
	}
}

func maiorNumero() {
	var numero int
	maior := 0

	for {
		fmt.Scan(&numero)
		if numero == 0 {
			break
		}
		if numero > maior {
			maior = numero
		}
	}
	fmt.Println(maior)

}

func buscaMatriz(m [][]int, n int, k int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if m[i][j] == k {
				return true
			}
		}

	}
	return false
}
func main() {
	//tomadas()

	m := [][]int{
		{1, 2, 3},
		{6, 4, 5},
		{7, 8, 9},
	}
	n := 3
	k := 0
	resultado := buscaMatriz(m, n, k)

	println(resultado)

}
