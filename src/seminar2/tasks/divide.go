package tasks

import "errors"

var ErrDivisionByZero = errors.New("division by zero")

// Divide выполняет деление двух чисел с обработкой ошибок
func Divide(a, b float64) (float64, error) {
	if b != 0 {
		return a / b, nil
	}
	return 0.0, ErrDivisionByZero
}
