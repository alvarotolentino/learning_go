package main

import "fmt"

func main() {
	s := foo()
	fmt.Println(s)

	x := bar()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)
}

func foo() string {
	s := "Hello world"
	return s
}

func bar() func() int {
	return func() int { return 100 }
}
