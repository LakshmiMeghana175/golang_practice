package main

import (
	"fmt"
)

func main() {
	aa:= []int{5,3,8,1,35,23}
	

	fmt.Println("Given array",a)
	//%%%%%%%% copy content of one array to another in reverse

	b := []int{}

	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}
	fmt.Println("In reverse order", b)
  }
