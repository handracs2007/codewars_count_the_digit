// https://www.codewars.com/kata/566fc12495810954b1000030/train/go

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func NbDig(n int, d int) int {
	count := 0

	for i := 0; i <= n; i++ {
		squared := i * i
		squaredStr := strconv.Itoa(squared)

		toCount := strconv.Itoa(d)

		count += strings.Count(squaredStr, toCount)
	}

	return count
}

func main() {
	fmt.Println(NbDig(550, 5))
	fmt.Println(NbDig(5750, 0))
}
