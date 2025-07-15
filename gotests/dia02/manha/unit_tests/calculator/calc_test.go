package calculator_test

import (
	"testing"

	"github.com/bgw7/unit_tests/calculator"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	// Dado
	num1 := 3
	num2 := 5
	expectedResult := 8

	// Quando
	result, err := calculator.Add(num1, num2)

	// Então
	require.Nil(t, err)
	require.Equal(t, expectedResult, result, "must be equal")

}

func TestSub(t *testing.T) {
	// Dado
	num1 := 5
	num2 := 3
	expectedResult := 2

	// Quando
	result, err := calculator.Sub(num1, num2)

	// Então
	require.Nil(t, err)
	require.Equal(t, expectedResult, result, "must be equal")

}
