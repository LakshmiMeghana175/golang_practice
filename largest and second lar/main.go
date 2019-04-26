package main

import (
	"fmt"
)

func main() {
	a := []int{5, 3, 8, 1, 35, 23}
	f := 0
	s := 0
	for _, ele := range a {
		if f < ele {
			f = ele

		} else if ele < f && ele > s {
			s = ele
		}
	}
	fmt.Println("Largest",f," Second largest",s)

}
