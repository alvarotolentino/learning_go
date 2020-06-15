package main

import "fmt"

func main() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains")
	case "swimming":
		fmt.Println("go to the pool")
	case "surfing":
		fmt.Println("go to Hawaii")
	case "wingsuit flying":
		fmt.Println("lol")
	}
}
