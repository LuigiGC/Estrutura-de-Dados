package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue []int

func (q *Queue) Push(n int) {
	*q = append(*q, n)
}

func (q *Queue) Pop() int {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		if n == 0 {
			break
		}

		queue := make(Queue, n)
		for i := 0; i < n; i++ {
			queue[i] = i + 1
		}

		discarded := []string{}
		for len(queue) >= 2 {
			discarded = append(discarded, strconv.Itoa(queue.Pop()))
			queue.Push(queue.Pop())
		}
		fmt.Printf("Discarded cards: ")
		fmt.Println(strings.Join(discarded, ", "))
		fmt.Printf("Remaining card: %d\n", queue[0])
	}
}
