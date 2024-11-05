package sample

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSample tests sample object
func TestSample(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "test1",
			expected: "test1",
		},
		{
			name:     "test2",
			expected: "test2",
		},
		{
			name:     "test3",
			expected: "test3",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSample(tc.name)
			assert.Equal(t, tc.expected, s.GetName())
		})
	}
}
