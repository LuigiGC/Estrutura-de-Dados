package main

import "fmt"

func main() {
	var casosTeste int
	fmt.Scanln(&casosTeste)

	for i := 0; i < casosTeste; i++ {
		var texto string
		var pos int
		fmt.Scanln(&texto)
		fmt.Scanln(&pos)
		var textoDecifrado string
		for _, letra := range texto {
				novaLetra := letra - rune(pos)
				if novaLetra < 'A' {
					novaLetra += 26
				}
				textoDecifrado += string(novaLetra)
		}
	fmt.Println(textoDecifrado)
	}
}
