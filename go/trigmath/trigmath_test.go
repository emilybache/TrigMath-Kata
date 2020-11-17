package trigmath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSin(t *testing.T) {
	var expected = float64(42.0)
	var trigmath = NewTrigMath()
	var actual = trigmath.Sin(3.4)
	assert.InDelta(t, expected, actual, 0.0000001)
}
