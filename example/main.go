package main

import (
	"fmt"

	regen "github.com/billbliu/goregen"
)

func main() {
	str, err := regen.Generate(`^[a-zA-Z0-9]{10,12}`)
	if err != nil {
		panic(err)
	}
	fmt.Println("str", str)
}
