package main

import (
	"fmt"
)

func sum(a []int) int {
	s := 0
	for _, i := range a {
		s += i
	}
	return s
}
func avg(a []int) float32 {

	return float32(sum(a) / len(a))
}

func main() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(sum(a))
	fmt.Println(avg(a))
}
