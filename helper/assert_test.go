package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Entong")
	assert.Equal(t, "Hello Fajar", result, "Harusnya Hello Fajar")
}
