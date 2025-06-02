package main

import "errors"

func division(numerator, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, errors.New("can not divide by zero")
	}
	return numerator / denominator, nil
}

func main() {
	result, err := division(10.0, 0.0)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Result:", result)
	}

	result, err = division(10.0, 2.0)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Result:", result)
	}
}
