// teste de caixa preta (black box)
// package calculator_test
// teste de caixa branca
package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {

	// opcao 1
	// t.Run("", func(t *testing.T) {})

	// opcao 2
	// table test
	// tableTest := []struct {
	// 	args struct {
	// 		num1 int64
	// 		num2 int64
	// 	}
	// 	want int64
	// }{
	// 	{
	// 		args: struct {
	// 			num1 int64
	// 			num2 int64
	// 		}{1, 2},
	// 		want: 3,
	// 	},
	// }

	// given (dado que) (preparações para o teste ou premissas)
	// define os parametros de teste
	// inicializa as instancias
	// valores esperados
	// inicializar os mocks
	num1 := 3
	num2 := 5
	expectedResult := 8

	// when (quando)
	result := Addition(num1, num2)

	// then (entao)
	if result != expectedResult {
		t.Errorf("Addition() function gave the result = %v, but the expected result is %v", result, expectedResult)
	}
}

func TestSubstraction(t *testing.T) {
	// given
	num1 := 3
	num2 := 5
	expectedResult := -2

	// when (quando)
	result := Substraction(num1, num2)

	// then (entao)
	require.Equal(t, expectedResult, result)
	// require.Nil(t, err)
	// require.NotNil(t, err)
	// require.Equal(t, "expected message error", err.Error())
}
