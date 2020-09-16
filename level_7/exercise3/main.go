package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4,8,3,32,99,18,12,45,23}
	xs := []string{"James", "George", "Thomas", "William", "Moneypenny"}

	fmt.Println(xi)
	fmt.Println(xs)
	
	sort.Strings(xs)
	sort.Ints(xi)
	
	fmt.Println(xi)
	fmt.Println(xs)

	

}