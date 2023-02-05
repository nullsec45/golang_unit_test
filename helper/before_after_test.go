package helper

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()
	fmt.Println("AFTER UNIT TEST")
}
func TestHelloWorldRequire2(t *testing.T) {
	result := HelloWorld("Entong")
	require.Equal(t, "Hello Fajar", result, "Harusnya Hello Fajar")
	fmt.Println("Ini tidak akan dieksekusi lagi")
}
