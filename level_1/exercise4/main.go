package main

import "fmt"

type (
	custom int
)

var x custom

func main() {

	x = 1

	fmt.Println("Value")
	fmt.Printf("\t%v\n", x)
	fmt.Println("Type")
	fmt.Printf("\t%T\n", x)
	fmt.Println("Memory Address")
	fmt.Printf("\t%v\n", &x)
}