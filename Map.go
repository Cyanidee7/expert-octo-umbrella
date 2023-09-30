package main

import (
	"fmt"
)

func main() {

	shop := map[string]int{
		"vandal":   2900,
		"guardian": 2250,
		"sheriff":  800,
		"marshal":  950,
	}

	fmt.Println(shop)
	fmt.Println("you bought the gun for", shop["sheriff"])
}
