package division

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntDivision(t *testing.T) {
	var tests = []struct {
		input1   int64
		input2   int64
		expected int64
	}{
		{200, 2, 100},
		{8, 2, 4},
		{20, 2, 10},
	}

	for _, test := range tests {
		res := IntDivision(test.input1, test.input2)
		assert.Equal(t, test.expected, res)
	}
}
