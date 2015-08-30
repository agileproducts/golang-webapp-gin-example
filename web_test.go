package main

import (
	"github.com/stretchr/testify/assert"
	_ "net/http"
	_ "net/http/httptest"
	"testing"
)

/* func TestSayHello(test *testing.T) {

	//it return Hello World on the homepage
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	assert.Equal(test, response.Body.String(), "Fuck You")

} */

func TestWhatever(test *testing.T) {
	x := Whatever()
	assert.Equal(test, x, "bob")
}
