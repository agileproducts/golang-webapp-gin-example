package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSayHello(test *testing.T) {

	router := gin.New()
	router.GET("/", SayHello)

	request, _ := http.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(test, "Hello World", writer.Body.String())

}

func TestSendSomeJson(test *testing.T) {

	router := gin.New()
	router.GET("/somejson", SomeJSON)

	type Msg struct {
		Key string `json:"key"`
	}

	var expected Msg
	expected.Key = "value"
	//expectedJson, _ := json.Marshal(expected)

	request, _ := http.NewRequest("GET", "/somejson", nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	actualBytes, _ := ioutil.ReadAll(writer.Body)
	actualMsg := Msg{}
	error := json.Unmarshal(actualBytes, &actualMsg)

	if error != nil {
		fmt.Println(error)
	}

	//Some debugging lines
	//fmt.Println(string(expectedJson))
	//fmt.Println(expectedJson)
	//fmt.Println(actualBytes)
	//fmt.Println(actualMsg)

	assert.Equal(test, expected, actualMsg)

}

func TestWhatever(test *testing.T) {
	x := whatever()
	assert.Equal(test, "bob", x)
}
