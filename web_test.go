package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"fmt"
	"bytes"
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
	router.GET("/somejson", SomeJson)

	type Msg struct {
		key	string
	}

	var expected Msg
	expected.key = "value"

	request, _ := http.NewRequest("GET", "/somejson", nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	fmt.Printf("%T\n",writer)
	fmt.Printf("%T\n",writer.Body)
	fmt.Printf("%T\n",writer.Body.Bytes())
	bob := *bytes.NewReader(writer.Body.Bytes())
	fmt.Printf("%T\n",bob)
	decoder := json.NewDecoder(bob)

	//http://stackoverflow.com/questions/25172177/json-marshal-how-body-of-http-newrequest


	//This is stupid. Json Unmarshal takes a byte[] doesn't work with a stream
	//Decoder works with a stream but won't accept a byte[] as an input
	//In ruby this would take 0.02 seconds

  //decoder := json.NewDecoder(writer.Body.Bytes())
  //jsonData := Msg{}
  //actual := decoder.Decode(&jsonData)
  //actual := json.Unmarshal(writer.Body.Bytes(),&jsonData)
	//assert.Equal(test, expected, actual)
}

func TestWhatever(test *testing.T) {
	x := whatever()
	assert.Equal(test, "bob", x)
}

  //http://stackoverflow.com/questions/16634582/sending-json-with-go

