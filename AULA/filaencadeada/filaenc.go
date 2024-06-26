package main

import "fmt"

type No struct {
	valor int
	prox  *No
}

type Fila struct {
	tamanho int
	topo  *No
	fim     *No
}
func (f *Fila) percorre() {
	no := f.topo
	if f.topo == nil {
		fmt.Println("Fila vazia!")
	} else {
		for no.prox != nil {
			fmt.Print(no.valor, " -> ")
			no = no.prox
		}
		fmt.Println(no.valor)
	}
}
func (f *Fila) inserir(novoValor int) {
	no := &No{valor: novoValor}
	if f.topo == nil {
		f.topo = no
	} else {
		no_prox := f.topo
		for no_prox.prox != nil {
			no_prox = no_prox.prox
		}
		no_prox.prox = no
	}

	f.tamanho++
}
func (f *Fila) remove() error {
	if f.topo == nil {
		return fmt.Errorf("Fila vazia")
	}
	aux := f.topo
	f.topo = aux.prox
	aux.prox = nil

	if f.topo == nil {
		f.topo = nil
	}
	f.tamanho--
	return nil
}


func main(){
	var fila Fila
	fila.inserir(2)
	fila.inserir(4)
	fila.inserir(8)

	fila.percorre()
	fmt.Println(fila.tamanho)

	fila.remove()

	fila.percorre()
	fmt.Println(fila.tamanho)

	fila.remove()
	fila.remove()

	fila.percorre()
	fmt.Println(fila.tamanho)

	err := fila.remove()
	fmt.Println(err)
}