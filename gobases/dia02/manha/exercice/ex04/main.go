// Os professores de uma universidade na Colômbia precisam
// calcular algumas estatísticas de notas para os alunos de um
// curso. Para isso, eles precisam gerar uma função que indique
// o tipo de cálculo que desejam realizar (mínimo, máximo ou médio)
// e que retorne outra função e uma mensagem (caso o cálculo não
// esteja definido) que possa receber um número N de inteiros e
// retorne o cálculo indicado na função anterior. Exemplo:

// const (
//    minimum = "minimum"
//    average = "average"
//    maximum = "maximum"
// )

// ...

// minFunc, err := operation(minimum)
// averageFunc, err := operation(average)
// maxFunc, err := operation(maximum)

// ...

// minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
// averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
// maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minValue(numbers ...float64) float64 {
	minValue := numbers[0]
	for _, number := range numbers {
		minValue = math.Min(minValue, number)
	}

	return minValue
}

func maxValue(numbers ...float64) float64 {
	maxValue := numbers[0]
	for _, number := range numbers {
		maxValue = math.Max(maxValue, number)
	}

	return maxValue
}

func averageValue(numbers ...float64) float64 {
	var sum float64

	for _, number := range numbers {
		sum += number
	}

	return sum / float64(len(numbers))
}

func operation(operation string) (func(numbers ...float64) float64, error) {
	switch operation {
	case "minimum":
		return minValue, nil
	case "maximum":
		return maxValue, nil
	case "average":
		return averageValue, nil
	default:
		return nil, errors.New("operation invalid")
	}
}

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		panic(err)
	}
	averageFunc, err := operation(average)
	if err != nil {
		panic(err)
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		panic(err)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minValue)
	fmt.Println(averageValue)
	fmt.Println(maxValue)
}
