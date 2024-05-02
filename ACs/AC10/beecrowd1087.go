package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minMovimentosDama(x1, y1, x2, y2 int) int {
	if x1 == x2 && y1 == y2 {
		return 0
	} else if x1 == x2 || y1 == y2 || abs(x1-x2) == abs(y1-y2) {
		return 1
	} else {
		return 2
	}
}

func main() {
	var x1, y1, x2, y2 int

	for {
		fmt.Scan(&x1, &y1, &x2, &y2)
		if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 {
			break
		}
		fmt.Println(minMovimentosDama(x1, y1, x2, y2))
	}
}