package main

import (
	"fmt"
)

func main() {

	var K int
	fmt.Scanln(&K)
	if K == 1 {
		fmt.Println("Top 1")
	} else if K <= 3 {
		fmt.Println("Top 3")
	} else if K <= 5 {
		fmt.Println("Top 5")
	} else if K <= 10 {
		fmt.Println("Top 10")
	} else if K <= 25 {
		fmt.Println("Top 25")
	} else if K <= 50 {
		fmt.Println("Top 50")
	} else if K <= 100 {
		fmt.Println("Top 100")
	}
}
