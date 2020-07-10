package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 7, 8, 9, 10}
	s1 := sum(ii...)

	fmt.Println("all numbers", s1)

	s2 := even(sum, ii...)
	fmt.Println("even number", s2)

	s3 := odd(sum, ii...)
	fmt.Println("odd number", s3)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, yi ...int) int {
	var zi []int
	for _, v := range yi {
		if v%2 == 0 {
			zi = append(zi, v)
		}
	}
	return f(zi...)
}

func odd(f func(xi ...int) int, yi ...int) int {
	var zi []int
	for _, v := range yi {
		if v%2 != 0 {
			zi = append(zi, v)
		}
	}
	return f(zi...)
}
