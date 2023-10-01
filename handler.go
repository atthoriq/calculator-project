package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	add      = "add"
	subtract = "subtract"
	multiply = "multiply"
	divide   = "divide"
	neg      = "neg"
	abs      = "abs"
	sqrt     = "sqrt"
	cbrt     = "cbrt"
	sqr      = "sqr"
	cube     = "cube"
	repeat   = "repeat"
	cancel   = "cancel"
	exit     = "exit"
	help     = "help"

	manual = `calculator will calculate new value to the current value. initial value will be 0.
add <float>      : add <float> to current
subtract <float> : subtract <float> to current
multiply <float> : add <float> to current
divide <float>   : add <float> to current
neg              : make current to negative. equally multiplying -1 to current. it requires no <float>
abs              : make current to positive. it requires no <float>
sqrt             : compute sqrt of current
cbrt             : compute cbrt of current
sqr              : compute sqr of current
cube             : compute cube of current
repeat <float>   : repeating <float> steps behind
cancel           : cancel calculation which set the current to 0.
exit             : exit the calculator
help             : show the manual`
)

// the interface is intended to mock the calculator easier
type Calculator interface {
	AddCurrent(a float64) float64
	SubtractCurrent(a float64) float64
	MultiplyCurrent(a float64) float64
	DivideCurrent(a float64) float64
	NegCurrent() float64
	AbsCurrent() float64
	RootCurrent(a float64) float64
	PowCurrent(a float64) float64
	Repeat(a float64) (float64, error)
	Cancel() float64
}

type calculatorHandler struct {
	calculator Calculator
}

func InitCalculatorHandler(calc Calculator) *calculatorHandler {
	return &calculatorHandler{
		calculator: calc,
	}
}

// Handle is to handle command string given from user
// to make no confusion, any commands requires only 1 argument will return error if they're given 2 or more
func (ch *calculatorHandler) Handle(command string) (string, error) {
	// sanitize leading and trailing white spaces
	command = strings.TrimSpace(command)

	commands := strings.Fields(command)
	if len(commands) > 2 || len(commands) == 0 {
		return "", errors.New("invalid input: read manual with 'help' command")
	}

	op := commands[0]

	var value float64
	if len(commands) == 2 {
		valueInStr := commands[1]
		v, err := strconv.ParseFloat(valueInStr, 64)
		if err != nil {
			return "", errors.New("invalid input: read manual with 'help' command")
		}
		value = v
	}

	switch op {
	case add:
		res := ch.calculator.AddCurrent(value)
		return fmt.Sprintf("%.2f", res), nil
	case subtract:
		res := ch.calculator.SubtractCurrent(value)
		return fmt.Sprintf("%.2f", res), nil
	case multiply:
		res := ch.calculator.MultiplyCurrent(value)
		return fmt.Sprintf("%.2f", res), nil
	case divide:
		res := ch.calculator.DivideCurrent(value)
		return fmt.Sprintf("%.2f", res), nil
	case neg:
		if len(commands) > 1 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		res := ch.calculator.NegCurrent()
		return fmt.Sprintf("%.2f", res), nil
	case abs:
		if len(commands) > 1 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		res := ch.calculator.AbsCurrent()
		return fmt.Sprintf("%.2f", res), nil
	case sqrt:
		if len(commands) > 1 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		res := ch.calculator.RootCurrent(2)
		return fmt.Sprintf("%.2f", res), nil
	case cbrt:
		if len(commands) > 1 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		res := ch.calculator.RootCurrent(3)
		return fmt.Sprintf("%.2f", res), nil
	case sqr:
		if len(commands) > 1 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		res := ch.calculator.PowCurrent(2)
		return fmt.Sprintf("%.2f", res), nil
	case cube:
		if len(commands) > 1 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		res := ch.calculator.PowCurrent(3)
		return fmt.Sprintf("%.2f", res), nil
	case repeat:
		res, err := ch.calculator.Repeat(value)
		return fmt.Sprintf("%.2f", res), err
	case cancel:
		res := ch.calculator.Cancel()
		return fmt.Sprintf("%.2f", res), nil
	case exit:
		if len(commands) > 1 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		return "", nil
	case help:
		if len(commands) > 1 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		return manual, nil
	default:
		return "", errors.New("not supported operation")
	}
}
