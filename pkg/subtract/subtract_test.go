package subtract

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSubtract(t *testing.T) {
	var tests = []struct {
		input1   int64
		input2   int64
		expected int64
	}{
		{10, 5, 5},
		{0, 0, 0},
		{-5, 2, -7},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, IntSubtract(test.input1, test.input2))
	}
}
