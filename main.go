package main

import (
	"fmt"
	"pingencli/pingen"
)

func main() {
	pin, err := pingen.GeneratePIN(6)
	if err != nil {
		panic(err)
	}

	fmt.Println(pin)
}
