package main

import (
	"fmt"
	m "projeto/modelos"
	"projeto/modelos/academico"
)

func main() {
	fmt.Println("Ola, mundo")

	aluno := m.Aluno{}
	aluno.Nome = "abcd"
	aluno.Matricula = "1234"

	curso := academico.Curso{Nome: "Engenharia"}

	fmt.Println(aluno, curso)
}
