package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSayHello(test *testing.T) {

	//it return Hello World on the homepage
	router := gin.New()
  router.GET("/", SayHello)

	request, _ := http.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(test, "Hello World", writer.Body.String())

}

func TestWhatever(test *testing.T) {
	x := whatever()
	assert.Equal(test, "bob", x)
}
