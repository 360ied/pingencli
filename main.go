package main

import (
	"fmt"
	"os"
	"pingencli/pingen"
	"strconv"
)

func main() {
	length := int64(6)
	if len(os.Args) > 1 {
		var err error
		length, err = strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			panic(err)
		}
	}

	pin, err := pingen.GeneratePIN(length)
	if err != nil {
		panic(err)
	}

	fmt.Println(pin)
}
