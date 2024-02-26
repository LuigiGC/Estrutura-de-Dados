package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}
func adicionarContato(contato Contato, arrayContatos []*Contato) {
	for i, c := range arrayContatos {
		if c == nil {
			arrayContatos[i] = &contato
			fmt.Println("Contato adicionado!")
			return
		}
	}
	fmt.Println("Não foi possível adicionar o contato pois já foram adicionados 5 contatos.")
}

func excluirContato(arrayContatos []*Contato) {
	for i := 4; i >= 0; i-- {
		if arrayContatos[i] != nil {
			arrayContatos[i] = nil
			fmt.Println("O último contato foi removido com sucesso.")
			return
		}
	}

	fmt.Println("Nenhum contato foi cadastrado.")
}

func main() {
	arrayContatos := make([]*Contato, 5)
	for {

		fmt.Println("\nEscolha uma opção:")
		fmt.Println("[1] Adicionar Contato")
		fmt.Println("[2] Excluir Último Contato")
		fmt.Println("[3] Sair")

		reader := bufio.NewReader(os.Stdin)
		escolha, _ := reader.ReadString('\n')
		escolha = strings.TrimSpace(escolha)

		switch escolha {
		case "1":
			var nome, email string
			fmt.Print("Nome do contato: ")
			fmt.Scanln(&nome)
			fmt.Print("E-mail do contato: ")
			fmt.Scanln(&email)
			novoContato := Contato{Nome: nome, Email: email}
			adicionarContato(novoContato, arrayContatos)
		case "2":
			excluirContato(arrayContatos)
		case "3":
			fmt.Println("Programa Encerrando.")
			return

		}
		fmt.Println("\nContatos:")
		for _, c := range arrayContatos {
			if c != nil {
				fmt.Printf("Nome: %s, Email: %s\n", c.Nome, c.Email)
			}

		}

	}
}
