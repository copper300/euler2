package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 1
	sum := 0
	for b < 4000000 {
		a, b = b, a+b
		if b%2 == 0 {
			sum += b
		}
	}
	fmt.Println("SUM is", sum)
}
