 
package main

import "fmt"

var x int = 3
var y string = "James Bond"
var z bool = true

func main() {
	// printing values
	fmt.Println("Values")
	s := fmt.Sprintf("\t:%v\t:%v\t:%v", x,y,z)
	fmt.Println(s)
	
	// printing typeof
	fmt.Println("Types")
	s = fmt.Sprintf("\t:%T\t:%T\t:%T", x, y, z)
	fmt.Println(s)
	
	// printing marmory address
	fmt.Println("Memory address")
	s = fmt.Sprintf("\t:%v\t:%v\t:%v", &x, &y, &z)
	fmt.Println(s)

}
