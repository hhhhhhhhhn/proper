# Proper
[Property Testing](https://www.youtube.com/watch?v=IYzDFHx6QPY) in Go, using reflection.

In essence, proper feeds random inputs (using [gofakeit](https://github.com/brianvoe/gofakeit))
to the properties, which are functions with arbitrary parameters
that return a boolean indicating if they are true or not.

See the [tests](./tests) for examples.

## Usage
```go
package tests

import (
	"testing"
	"github.com/hhhhhhhhhn/proper"
)

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
```
