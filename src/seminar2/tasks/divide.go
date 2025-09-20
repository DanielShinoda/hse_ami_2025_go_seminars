package tasks

import "errors"

// Divide выполняет деление двух чисел с обработкой ошибок
func Divide(a, b float64) (float64, error) {
	if b != 0 {
		return a / b, nil
	} else {
		return 0.0, errors.New("division by zero")
	}
}
