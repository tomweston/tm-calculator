package add

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntAdd(t *testing.T) {
	var tests = []struct {
		input1   int64
		input2   int64
		expected int64
	}{
		{1, 1, 2},
		{0, 0, 0},
		{-50, 30, -20},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, IntAdd(test.input1, test.input2))
	}
}
