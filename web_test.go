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
	request, _ := http.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(test, "Fuck You", writer.Body.String())

}

func TestWhatever(test *testing.T) {
	x := Whatever()
	assert.Equal(test, "bob", x)
}
