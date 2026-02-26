package calculator

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		operands []float64
		want     float64
		wantErr  bool
	}{
		{"two positive numbers", []float64{5, 3}, 8, false},
		{"multiple numbers", []float64{1, 2, 3, 4}, 10, false},
		{"negative numbers", []float64{-5, -3}, -8, false},
		{"mixed signs", []float64{10, -5}, 5, false},
		{"decimals", []float64{1.5, 2.5}, 4.0, false},
		{"single operand", []float64{5}, 0, true},
		{"no operands", []float64{}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate("add", tt.operands)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate(add) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Calculate(add) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		operands []float64
		want     float64
		wantErr  bool
	}{
		{"two positive numbers", []float64{5, 3}, 2, false},
		{"multiple numbers", []float64{10, 2, 3}, 5, false},
		{"negative result", []float64{3, 5}, -2, false},
		{"decimals", []float64{5.5, 2.5}, 3.0, false},
		{"single operand", []float64{5}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate("subtract", tt.operands)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate(subtract) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Calculate(subtract) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		operands []float64
		want     float64
		wantErr  bool
	}{
		{"two positive numbers", []float64{5, 3}, 15, false},
		{"multiple numbers", []float64{2, 3, 4}, 24, false},
		{"with zero", []float64{5, 0}, 0, false},
		{"negative numbers", []float64{-5, 3}, -15, false},
		{"decimals", []float64{2.5, 4}, 10.0, false},
		{"single operand", []float64{5}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate("multiply", tt.operands)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate(multiply) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Calculate(multiply) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		operands []float64
		want     float64
		wantErr  bool
		errType  error
	}{
		{"two positive numbers", []float64{10, 2}, 5, false, nil},
		{"with decimals", []float64{7, 2}, 3.5, false, nil},
		{"negative divisor", []float64{10, -2}, -5, false, nil},
		{"division by zero", []float64{10, 0}, 0, true, ErrDivisionByZero},
		{"wrong operand count", []float64{10, 2, 3}, 0, true, nil},
		{"single operand", []float64{10}, 0, true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate("divide", tt.operands)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate(divide) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.errType != nil && err != tt.errType {
				t.Errorf("Calculate(divide) error = %v, want %v", err, tt.errType)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Calculate(divide) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExponentiation(t *testing.T) {
	tests := []struct {
		name     string
		operands []float64
		want     float64
		wantErr  bool
		errType  error
	}{
		{"positive base and exponent", []float64{2, 3}, 8, false, nil},
		{"base to zero power", []float64{5, 0}, 1, false, nil},
		{"negative exponent", []float64{2, -2}, 0.25, false, nil},
		{"fractional exponent", []float64{4, 0.5}, 2, false, nil},
		{"zero to positive power", []float64{0, 5}, 0, false, nil},
		{"zero to negative power", []float64{0, -1}, 0, true, ErrInvalidExponent},
		{"negative base fractional exp", []float64{-4, 0.5}, 0, true, ErrInvalidExponent},
		{"negative base integer exp", []float64{-2, 3}, -8, false, nil},
		{"wrong operand count", []float64{2}, 0, true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate("exponentiation", tt.operands)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate(exponentiation) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.errType != nil && err != tt.errType {
				t.Errorf("Calculate(exponentiation) error = %v, want %v", err, tt.errType)
			}
			if !tt.wantErr && math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("Calculate(exponentiation) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		name     string
		operands []float64
		want     float64
		wantErr  bool
		errType  error
	}{
		{"positive number", []float64{9}, 3, false, nil},
		{"zero", []float64{0}, 0, false, nil},
		{"decimal", []float64{2.25}, 1.5, false, nil},
		{"negative number", []float64{-4}, 0, true, ErrNegativeRoot},
		{"wrong operand count", []float64{4, 9}, 0, true, nil},
		{"no operands", []float64{}, 0, true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate("sqrt", tt.operands)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate(sqrt) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.errType != nil && err != tt.errType {
				t.Errorf("Calculate(sqrt) error = %v, want %v", err, tt.errType)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Calculate(sqrt) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercentage(t *testing.T) {
	tests := []struct {
		name     string
		operands []float64
		want     float64
		wantErr  bool
	}{
		{"simple percentage", []float64{100, 10}, 10, false},
		{"fractional percentage", []float64{50, 25}, 12.5, false},
		{"percentage over 100", []float64{100, 150}, 150, false},
		{"zero percentage", []float64{100, 0}, 0, false},
		{"wrong operand count", []float64{100}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate("percentage", tt.operands)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate(percentage) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Calculate(percentage) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvalidOperation(t *testing.T) {
	_, err := Calculate("invalid", []float64{1, 2})
	if err != ErrInvalidOperation {
		t.Errorf("Calculate(invalid) error = %v, want %v", err, ErrInvalidOperation)
	}
}
