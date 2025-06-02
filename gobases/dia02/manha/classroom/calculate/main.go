package main

import "fmt"

const (
	Sum  = "+"
	Sub  = "-"
	Mult = "*"
	Div  = "/"
)

func opSummatory(value1, value2 float64) float64 {
	return value1 + value2
}

func opSubtraction(value1, value2 float64) float64 {
	return value1 - value2
}

func opMultiplication(value1, value2 float64) float64 {
	return value1 * value2
}

func opDivision(value1, value2 float64) float64 {

	if value2 == 0 {
		return 0
	}
	return value1 / value2
}

func calculate(operator string, values ...float64) float64 {
	switch operator {
	case Sum:
		return operationsOrchestrator(values, opSummatory)
	case Sub:
		return operationsOrchestrator(values, opSubtraction)
	case Mult:
		return operationsOrchestrator(values, opMultiplication)
	case Div:
		return operationsOrchestrator(values, opDivision)
	}

	return 0
}
func operationsOrchestrator(
	values []float64,
	operation func(value1, value2 float64) float64,
) float64 {
	var result float64
	for _, value := range values {
		result = operation(result, value)
	}

	return result
}

func main() {
	resultSum := calculate(Sum, 1.0, 2.0, 3.0, 4.0, 5.0)
	resultSub := calculate(Sub, 1.0, 2.0, 3.0, 4.0, 5.0)
	fmt.Println("The result of the addition is:", resultSum)
	fmt.Println("The result of the sub is:", resultSub)
}
