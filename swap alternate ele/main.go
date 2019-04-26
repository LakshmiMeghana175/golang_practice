package main

import (
	"fmt"
)

func main() {
	a :=[]int{1,2,3,4,5,6,7,8,9}
	
	for i:=0; i<len(a)-1; i+=2 {
	 
	a[i],a[i+1]=a[i+1],a[i]
	
	
	}
	
 fmt.Println(a)

}
