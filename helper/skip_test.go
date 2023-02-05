package helper

import (
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Unit test tidak bisa jalan di windows")
	}
	result := HelloWorld("Entong")
	require.Equal(t, "Hello Fajar", result, "Harusnya Hello Fajar")
}
