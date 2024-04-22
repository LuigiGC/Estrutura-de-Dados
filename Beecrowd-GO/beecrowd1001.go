package main

import (
    "fmt"
)

func main() {

    var A int
    var B int
    var X int

    fmt.Scanln(&A)
    fmt.Scanln(&B)

    X = A + B
    fmt.Printf("X = %d\n",X)
}
