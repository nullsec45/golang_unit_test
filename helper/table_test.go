package helper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Rama",
			request:  "Rama",
			expected: "Hello Rama",
		},
		{
			name:     "Fajar",
			request:  "Fajar",
			expected: "Hello Fajar",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
