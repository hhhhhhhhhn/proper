package tests

import (
	"testing"
	"github.com/hhhhhhhhhn/proper"
)

// Wrong
func negate(x int16) int16 {
	return 2 * x
}

// Will actually work, as doubling is associative as well
func associative(x, y int16) bool {
	return negate(x + y) == negate(x) + negate(y)
}

// Will fail
func addsToZero(x int16) bool {
	return negate(x) + x == 0
}

func TestNegateProperties(t *testing.T) {
	proper.TestProperties(t, associative, addsToZero)
}
