package factorial

func Factorial(number int) int {
	// if number == 0 || number == 1 {
	// 	return 1
	// }
	// return number * Factorial(number-1)

	f := 1
	for i := 1; i <= number; i++ {
		f *= i
	}

	return f
}

// 2 => 2 x 1
// 3 => 3 * 2 * 1
