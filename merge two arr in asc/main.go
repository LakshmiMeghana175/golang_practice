package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0)
	a1 := []int{1, 2, 8, 10, 7}
	a2 := []int{4, 9, 5, 6, 3}
	a1 = sort(a1)
	a2 = sort(a2)

	j := 0
	k := 0
	for i := 0; i < len(a1)+len(a2); i++ {
		if a1[j] < a2[k] {

			if j < len(a1) {

				a = append(a, a1[j])
				j++

			}

			if j == len(a1) {
				for k < len(a2) {
					a = append(a, a2[k])
					k++
					i++
				}
			}

		} else if a1[j] > a2[k] {
			if k < len(a2) {

				a = append(a, a2[k])
				k++

			}

			if k == len(a2) {
				for j < len(a1) {
					a = append(a, a1[j])
					j++
					i++
				}
			}
		}

	}
	fmt.Println(a)
}

func sort(t []int) []int {

	for i := 0; i < len(t); i++ {
		for j := i + 1; j < len(t); j++ {

			if t[j] < t[i] {

				t[i], t[j] = t[j], t[i]
			}
		}
	}
	return t

}
