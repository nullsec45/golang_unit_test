package helper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSubTest(t *testing.T) {
	t.Run("Fajar", func(t *testing.T) {
		result := HelloWorld("Fajar")
		require.Equal(t, "Hello Fajar", result, "Harusnya Hello Fajar")
	})
	t.Run("Entong", func(t *testing.T) {
		result := HelloWorld("Entong")
		require.Equal(t, "Hello Entong", result, "Harusnya Hello Entong")
	})
}
