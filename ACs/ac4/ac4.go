package main

import "fmt"

func honoi(n int, origem, destino, aux string) {
	if n == 1 {
		fmt.Println("Mova o disco de cima de ",origem," Para ",destino)
		return
	}
	honoi(n-1,origem,aux,destino)
	fmt.Println("Mova o disco ",n," de ",origem," Para", destino)
	honoi(n-1, aux,destino,origem)
}

func main() {
 discos := 4
 honoi(discos, "Torre 1","Torre 3", "Torre 2")
}
