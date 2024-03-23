package main

import (
	"fmt"
)

type Fracao struct {
	num int
	den int
}

func soma(f1, f2 Fracao) Fracao {
	retorno := Fracao{
		num: f1.num*f2.den + f2.num*f1.den,
		den: f1.den * f2.den,
	}
	return retorno
}

func subtracao(f1, f2 Fracao) Fracao {
	retorno := Fracao{
		num: f1.num*f2.den - f2.num*f1.den,
		den: f1.den * f2.den,
	}
	return retorno
}

func multiplicacao(f1, f2 Fracao) Fracao {
	retorno := Fracao{
		num: f1.num * f2.num,
		den: f1.den * f2.den,
	}
	return retorno
}

func divisao(f1, f2 Fracao) Fracao {
	retorno := Fracao{
		num: f1.num * f2.den,
		den: f2.num * f1.den,
	}
	return retorno
}

func mdc(a, b int) int {
	if b == 0 {
		return a
	}
	return mdc(b, a%b)
}

func irredutivel(f Fracao) Fracao {
	mdc := mdc(abs(f.num), abs(f.den))
	f.num /= mdc
	f.den /= mdc
	return f
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var N int
	var op string
	var f1, f2, resultado Fracao

	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d / %d %s %d / %d", &f1.num, &f1.den, &op, &f2.num, &f2.den)

		switch op {
		case "+":
			resultado = soma(f1, f2)
		case "-":
			resultado = subtracao(f1, f2)
		case "*":
			resultado = multiplicacao(f1, f2)
		case "/":
			resultado = divisao(f1, f2)
		}

		fmt.Printf("%d/%d = ", resultado.num, resultado.den)
		resultado = irredutivel(resultado)
		fmt.Printf("%d/%d\n", resultado.num, resultado.den)
	}
}
