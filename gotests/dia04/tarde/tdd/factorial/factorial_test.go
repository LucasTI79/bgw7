package factorial_test

import (
	"testing"

	"github.com/lucasti79/bgw7/tdd/factorial"
	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	// given
	expectedResult := 6
	testArg := 3

	// when
	result := factorial.Factorial(testArg)

	// then
	assert.Equal(t, expectedResult, result)
}
