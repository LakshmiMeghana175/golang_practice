package main

import (
	"fmt"
)

func main() {
	a:= []int{5,3,8,1,35,23}
	

	// %%%% arrange in ascending order

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println("Ascending order",a)
	
	

}
