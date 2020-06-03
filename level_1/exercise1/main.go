package main

import "fmt"

func main() {
	x := 42.0
	y := "James Bond"
	z := true
	PrintValues(x, y, z)
	PrintTypeof(x)
	PrintTypeof(y)
	PrintTypeof(z)
	PrintAddress(x)
	PrintAddress(y)
	PrintAddress(z)

}

// PrintValues prints the values of all variables
func PrintValues(i ...interface{}) {
	fmt.Printf("%v \n", i)
}

// PrintValue prints the value of a variable
func PrintValue(i interface{}) {
	fmt.Printf("%v \n", i)
}

// PrintTypeof prints the typeof of a variable
func PrintTypeof(i interface{}) {
	fmt.Printf("%T \n", i)
}

// PrintAddress prints the memory address of the variable
func PrintAddress(i interface{}) {
	fmt.Printf("%v \n", &i)
}
