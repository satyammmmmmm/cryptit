package main

import (
	"fmt"

	"satyam/decrpt"
	"satyam/encypt"

	"github.com/gin-gonic/gin"
)

func main() {
	_ := gin.New()
	fmt.Println(encypt.Nimbus("satyam"))
	fmt.Println(decrpt.Nimbus())
}
