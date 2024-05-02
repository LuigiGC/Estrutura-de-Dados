package main

import "fmt"

func main() {
	var casosTeste int
	fmt.Scanln(&casosTeste)

	for i := 0; i < casosTeste; i++ {
		Tled := 0
		var led string
		fmt.Scanln(&led)

		for _, j := range led {
			if j == '1' {
				Tled += 2
			} else if j == '2' || j == '3' || j == '5' {
				Tled += 5
			} else if j == '4' {
				Tled += 4
			} else if j == '6' || j == '9' || j == '0' {
				Tled += 6
			} else if j == '7' {
				Tled += 3
			} else if j == '8' {
				Tled += 7
			}
		}
		fmt.Println(Tled, "leds")
	}
}
