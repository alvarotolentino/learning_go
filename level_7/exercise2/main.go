package main

import (
	"encoding/json"
	"fmt"
)

var jsonBlob = []byte(`[
	{"Manufactured": "Ford", "Model": "F150", "Doors": 2},
	{"Manufactured": "Toyota", "Model": "Yaris", "Doors": 4}
]`)

type car struct {
	Manufactured string
	Model        string
	Doors        int
}

func main() {
	var cars []car
	err := json.Unmarshal(jsonBlob, &cars)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("all of data:", cars)

	for i, car := range cars {
		fmt.Println("Car #", i)
		fmt.Println(car.Manufactured, car.Model, car.Doors)
	}
}
