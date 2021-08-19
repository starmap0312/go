package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 1) unmarshall a json string (byte array) to a struct
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata", "Price": 100},
	{"Name": "Quoll",    "Order": "Dasyuromorphia", "Price": 200}]`)

	type Animal struct {
		Name  string
		Order string
		Price int
	}

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err) // ex. error: json: cannot unmarshal number into Go struct field Animal.Price of type string
	}
	fmt.Println(animals) // [{Platypus Monotremata} {Quoll Dasyuromorphia}]

}
