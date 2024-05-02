package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {

        input := strings.Fields(scanner.Text())
        chave := input[0]
        texto := input[1]


        if chave == "0" && texto == "0" {
            break
        }
        textoModificado := strings.ReplaceAll(texto, chave, "")
        textoModificado = strings.TrimLeft(textoModificado, "0")

        if textoModificado == "" {
            textoModificado = "0"
        }       
        fmt.Println(textoModificado)
    }
}
