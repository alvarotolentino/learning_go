package main

import "fmt"

type (custom int)

var x custom
var y int

func main() {
	fmt.Println("Value")
	fmt.Printf("\t%v\n", x)
	fmt.Println("Type")
	fmt.Printf("\t%T\n", x)
	
	x = 4
	fmt.Println("After assign a value")
	fmt.Printf("\t%v\n", x)
	
	y =  int(x)
	fmt.Println("Value of y variable")
	fmt.Printf("\t%v\n", y)
	fmt.Println("Type of y variable")
	fmt.Printf("\t%T\n", y)
	

}