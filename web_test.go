package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(test *testing.T) {

	//it should assert equality
	assert.Equal(test, 123, 123, "they should be equal")

}
