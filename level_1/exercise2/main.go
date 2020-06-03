package main

import "fmt"

var x int
var y string
var z bool
func main() {
	// printing values
	fmt.Println("zero-values")
	fmt.Printf("\t:%v\t:%v\t:%v\n", x, y, z)
	
	// printing typeof
	fmt.Println("types")
	fmt.Printf("\t:%T\t:%T\t:%T\n", x, y, z)
	
	// printing marmory address
	fmt.Println("memory address")
	fmt.Printf("\t:%v\t:%v\t:%v\n", &x, &y, &z)
}
