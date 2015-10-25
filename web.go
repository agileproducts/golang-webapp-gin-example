package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/", SayHello)
	router.GET("/somejson", SomeJSON)

	router.Run(":5000")

}

//SayHello says hello
func SayHello(context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}

//SomeJSON returns some json
func SomeJSON(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"key": "value"})
}

func whatever() string {
	return "bob"
}
