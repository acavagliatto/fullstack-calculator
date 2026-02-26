package calculator

import (
	"errors"
	"math"
)

var (
	ErrDivisionByZero     = errors.New("division by zero")
	ErrInvalidOperation   = errors.New("invalid operation")
	ErrNegativeRoot       = errors.New("cannot compute square root of negative number")
	ErrInvalidExponent    = errors.New("invalid exponent operation")
)

// Calculate performs the specified operation on the given operands
func Calculate(operation string, operands []float64) (float64, error) {
	switch operation {
	case "add":
		return add(operands)
	case "subtract":
		return subtract(operands)
	case "multiply":
		return multiply(operands)
	case "divide":
		return divide(operands)
	case "exponentiation":
		return exponentiation(operands)
	case "sqrt":
		return sqrt(operands)
	case "percentage":
		return percentage(operands)
	default:
		return 0, ErrInvalidOperation
	}
}

func add(operands []float64) (float64, error) {
	if len(operands) < 2 {
		return 0, errors.New("add requires at least 2 operands")
	}
	result := operands[0]
	for i := 1; i < len(operands); i++ {
		result += operands[i]
	}
	return result, nil
}

func subtract(operands []float64) (float64, error) {
	if len(operands) < 2 {
		return 0, errors.New("subtract requires at least 2 operands")
	}
	result := operands[0]
	for i := 1; i < len(operands); i++ {
		result -= operands[i]
	}
	return result, nil
}

func multiply(operands []float64) (float64, error) {
	if len(operands) < 2 {
		return 0, errors.New("multiply requires at least 2 operands")
	}
	result := operands[0]
	for i := 1; i < len(operands); i++ {
		result *= operands[i]
	}
	return result, nil
}

func divide(operands []float64) (float64, error) {
	if len(operands) != 2 {
		return 0, errors.New("divide requires exactly 2 operands")
	}
	if operands[1] == 0 {
		return 0, ErrDivisionByZero
	}
	return operands[0] / operands[1], nil
}

func exponentiation(operands []float64) (float64, error) {
	if len(operands) != 2 {
		return 0, errors.New("exponentiation requires exactly 2 operands")
	}
	
	// Check for invalid cases
	if operands[0] == 0 && operands[1] < 0 {
		return 0, ErrInvalidExponent
	}
	if operands[0] < 0 && math.Floor(operands[1]) != operands[1] {
		return 0, ErrInvalidExponent
	}
	
	result := math.Pow(operands[0], operands[1])
	if math.IsInf(result, 0) || math.IsNaN(result) {
		return 0, ErrInvalidExponent
	}
	
	return result, nil
}

func sqrt(operands []float64) (float64, error) {
	if len(operands) != 1 {
		return 0, errors.New("sqrt requires exactly 1 operand")
	}
	if operands[0] < 0 {
		return 0, ErrNegativeRoot
	}
	return math.Sqrt(operands[0]), nil
}

func percentage(operands []float64) (float64, error) {
	if len(operands) != 2 {
		return 0, errors.New("percentage requires exactly 2 operands (value, percentage)")
	}
	return operands[0] * operands[1] / 100, nil
}
