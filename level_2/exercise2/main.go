package main

import "fmt"

func main() {
	a := (10 == 10)
	b := (11 <= 12)
	c := (12 >= 13)
	d := (13 != 14)
	e := (15 < 16)
	f := (16 > 17)

	fmt.Println(a,b,c,d,e,f)
}
