package main

import (
	"fmt"

	"satyam/decrpt"
	"satyam/encypt"
)

func main() {
	fmt.Println(encypt.Nimbus("satyam"))
	fmt.Println(decrpt.Nimbus())
}
