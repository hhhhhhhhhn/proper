package tests

import (
	"testing"
	"github.com/hhhhhhhhhn/proper"
)

// Correct
func add(x, y int16) int16 {
	return x + y
}

func commutative(x, y int16) bool {
	return add(x, y) == add(y, x)
}

func addTwiceIsAddingTwo(x int16) bool {
	return add(x, 2) == add(add(x, 1), 1)
}

func addingZeroDoesNothing(x int16) bool {
	return add(x, 0) == x
}

func TestAdditionProperties(t *testing.T) {
	proper.TestProperties(t, commutative, addTwiceIsAddingTwo, addingZeroDoesNothing)
}
