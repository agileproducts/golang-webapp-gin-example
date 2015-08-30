package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/", SayHello)

	router.Run(":5000")

}

func SayHello(context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}

func Whatever() string {
	return "bob"
}
