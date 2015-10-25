package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	router.GET("/", SayHello)
	router.GET("/somejson", SomeJSON)

	router.Run(":" + port)

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
