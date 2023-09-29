package main

import (
	"fmt"

	"github.com/satyammmmmmm/cryptit/encypt"

	"github.com/gin-gonic/gin"
)

func main() {
	_ = gin.New()
	fmt.Println(encypt.Nimbus("satyam"))

}
