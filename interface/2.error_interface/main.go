package main

import (
	"errors"
	"fmt"
)

var ErrDivisionByZero = errors.New("деление на 0")
var ErrNegativeNumber = errors.New("отрицательное число")

func SafeDivide(a, b int) (int, error) {
	res, err := Divide(a, b)
	if err != nil {
		return 0, fmt.Errorf("не удалось выполнить SafeDivide: %w", err)

	} else {
		return res, nil
	}

}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	} else if a < 0 || b < 0 {
		return 0, ErrNegativeNumber

	}
	return a / b, nil
}
func main() {
	a, b := 10, -1
	res, err := SafeDivide(a, b)
	if errors.Is(err, ErrDivisionByZero) {
		fmt.Println("обнаружено деление на 0!")

	} else if errors.Is(err, ErrNegativeNumber) {
		fmt.Println("обнаружено отрицательное число!")
	} else {
		fmt.Println(res)
	}

	c, d := 10, 0

	res2, err2 := SafeDivide(c, d)
	if errors.Is(err2, ErrDivisionByZero) {
		fmt.Println("обнаружено деление на 0!")

	} else if errors.Is(err2, ErrNegativeNumber) {
		fmt.Println("обнаружено отрицательное число!")
	} else {
		fmt.Println(res2)
	}

	e, f := 10, 2

	res3, err3 := SafeDivide(e, f)
	if errors.Is(err3, ErrDivisionByZero) {
		fmt.Println("обнаружено деление на 0!")

	} else if errors.Is(err3, ErrNegativeNumber) {
		fmt.Println("обнаружено отрицательное число!")
	} else {
		fmt.Println(res3)
	}
}
