package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	traducoes := map[string]string{
		"brasil":         "Feliz Natal!",
		"alemanha":       "Frohliche Weihnachten!",
		"austria":        "Frohe Weihnacht!",
		"coreia":         "Chuk Sung Tan!",
		"espanha":        "Feliz Navidad!",
		"grecia":         "Kala Christougena!",
		"estados-unidos": "Merry Christmas!",
		"inglaterra":     "Merry Christmas!",
		"australia":      "Merry Christmas!",
		"portugal":       "Feliz Natal!",
		"suecia":         "God Jul!",
		"turquia":        "Mutlu Noeller",
		"argentina":      "Feliz Navidad!",
		"chile":          "Feliz Navidad!",
		"mexico":         "Feliz Navidad!",
		"antardida":      "Merry Christmas!",
		"canada":         "Merry Christmas!",
		"irlanda":        "Nollaig Shona Dhuit!",
		"belgica":        "Zalig Kerstfeest!",
		"italia":         "Buon Natale!",
		"libia":          "Buon Natale!",
		"siria":          "Milad Mubarak!",
		"marrocos":       "Milad Mubarak!",
		"japao":          "Merii Kurisumasu!",
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		pais := strings.TrimSpace(scanner.Text())

		if pais == "" {
			break
		}
		if traducao, ok := traducoes[pais]; ok {
			fmt.Println(traducao)
		} else {
			fmt.Println("--- NOT FOUND ---")
		}
	}
}
