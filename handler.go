package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/atthoriq/calculator-project/calculator"
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

type calculatorHandler struct {
	calculator calculator.NewCalculator
}

func InitCalculatorHandler(calc calculator.NewCalculator) *calculatorHandler {
	return &calculatorHandler{
		calculator: calc,
	}
}

// Handle is to handle command string given from user
// to make no confusion, any commands requires only 1 argument will return error if they're given 2 or more
func (ch *calculatorHandler) Handle(command string) (string, error) {
	// sanitize leading and trailing white spaces
	op, value, err := parseCommand(command)
	if err != nil {
		return "", err
	}

	switch op {
	case add:
		res := ch.calculator.Add(value).GetResult()
		return fmt.Sprintf("%.2f", res), nil
	case subtract:
		res := ch.calculator.Subtract(value).GetResult()
		return fmt.Sprintf("%.2f", res), nil
	case multiply:
		res := ch.calculator.Multiply(value).GetResult()
		return fmt.Sprintf("%.2f", res), nil
	case divide:
		res := ch.calculator.Divide(value).GetResult()
		return fmt.Sprintf("%.2f", res), nil
	case neg:
		if value > 0 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		res := ch.calculator.Multiply(-1).GetResult()
		return fmt.Sprintf("%.2f", res), nil
	case abs:
		if value > 0 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		res := ch.calculator.Abs().GetResult()
		return fmt.Sprintf("%.2f", res), nil
	case sqrt:
		if value > 0 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		res := ch.calculator.Root(2).GetResult()
		return fmt.Sprintf("%.2f", res), nil
	case cbrt:
		if value > 0 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		res := ch.calculator.Root(3).GetResult()
		return fmt.Sprintf("%.2f", res), nil
	case sqr:
		if value > 0 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		res := ch.calculator.Pow(2).GetResult()
		return fmt.Sprintf("%.2f", res), nil
	case cube:
		if value > 0 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		res := ch.calculator.Pow(3).GetResult()
		return fmt.Sprintf("%.2f", res), nil
	case repeat:
		res := ch.calculator.Repeat(int(value)).GetResult()
		return fmt.Sprintf("%.2f", res), err
	case cancel:
		res := ch.calculator.Cancel().GetResult()
		return fmt.Sprintf("%.2f", res), nil
	case exit:
		if value > 0 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		return "", nil
	case help:
		if value > 0 {
			return "", errors.New("invalid input: read manual with 'help' command")
		}

		return manual, nil
	default:
		return "", errors.New("not supported operation")
	}
}

func parseCommand(command string) (op string, val float64, err error) {
	command = strings.TrimSpace(command)

	commands := strings.Fields(command)
	if len(commands) > 2 || len(commands) == 0 {
		return "", 0, errors.New("invalid input: read manual with 'help' command")
	}

	op = commands[0]
	if len(commands) == 2 {
		valueInStr := commands[1]
		v, err := strconv.ParseFloat(valueInStr, 64)
		if err != nil {
			return "", 0, errors.New("invalid input: read manual with 'help' command")
		}
		val = v
	}

	return op, val, nil
}
