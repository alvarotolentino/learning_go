package main

import "fmt"

func main() {
	fmt.Println("Example of func expression")

	f := func() {
		fmt.Println("Call a func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("Call a func expression with parameter")
	}
	g(39)
}
