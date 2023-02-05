package helper

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Entong")
	require.Equal(t, "Hello Fajar", result, "Harusnya Hello Fajar")
	fmt.Println("Ini tidak akan dieksekusi lagi")
}
