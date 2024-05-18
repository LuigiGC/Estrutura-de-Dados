package main

import "fmt"

func main() {
	for {
		var H1, M1, H2, M2 int
		var tempo int

		_, err := fmt.Scanf("%d %d %d %d", &H1, &M1, &H2, &M2)
		if err != nil {
			continue
		}

		if H1 == 0 && M1 == 0 && H2 == 0 && M2 == 0 {
			break
		}
		horasMinutos := H1*60 + M1
		alarmeHorasMinutos := H2*60 + M2
		if alarmeHorasMinutos >= horasMinutos {
			tempo = alarmeHorasMinutos - horasMinutos
		} else {
			tempo = (24*60 - horasMinutos) + alarmeHorasMinutos
		}
		fmt.Println(tempo)
	}

}
