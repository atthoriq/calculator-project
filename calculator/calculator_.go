package calculator

import (
	"errors"
	"fmt"
	"math"
)

const (
	addOp      = "add"
	subtractOp = "subtract"
	multiplyOp = "multiply"
	divideOp   = "divide"
	negOp      = "neg"
	rootOp     = "root"
	powOp      = "pow"
	cancelOp   = "cancel"
	absOp      = "abs"
	repeatOp   = "repeat"
)

type Calculator struct {
	history []*command
	current float64
}

type command struct {
	op   string // its a function
	args []float64
}

func InitCalculator() *Calculator {
	c := &Calculator{
		current: 0,
		history: make([]*command, 0),
	}
	return c
}

func (c *Calculator) Add(a float64) float64 {
	c.addHistory(addOp, []float64{a})
	return c.addOp(a)
}

func (c *Calculator) addOp(a float64) float64 {
	c.current = c.current + a
	return c.current
}

func (c *Calculator) Subtract(a float64) float64 {
	c.addHistory(subtractOp, []float64{a})
	return c.subtractOp(a)
}

func (c *Calculator) subtractOp(a float64) float64 {
	c.current = c.current - a
	return c.current
}

func (c *Calculator) Multiply(a float64) float64 {
	c.addHistory(multiplyOp, []float64{a})
	return c.multiplyOp(a)
}

func (c *Calculator) multiplyOp(a float64) float64 {
	c.current = c.current * a
	return c.current
}

func (c *Calculator) Divide(a float64) float64 {
	c.addHistory(divideOp, []float64{a})
	return c.divideOp(a)
}

func (c *Calculator) divideOp(a float64) float64 {
	if a == 0 {
		c.current = math.NaN()
	} else {
		c.current = c.current / a
	}

	return c.current
}

func (c *Calculator) Neg() float64 {
	c.addHistory(negOp, []float64{})
	return c.negOp()
}

func (c *Calculator) negOp() float64 {
	c.current = c.multiplyOp(-1)
	return c.current
}

func (c *Calculator) Abs() float64 {
	c.addHistory(absOp, []float64{})
	return c.absOp()
}

func (c *Calculator) absOp() float64 {
	c.current = math.Abs(c.current)
	return c.current
}

func (c *Calculator) Root(n float64) float64 {
	c.addHistory(rootOp, []float64{n})
	return c.rootOp(n)
}

func (c *Calculator) rootOp(n float64) float64 {
	switch n {
	case 2:
		c.current = math.Sqrt(c.current)
		return c.current
	case 3:
		c.current = math.Cbrt(c.current)
		return c.current
	default:
		return math.NaN()
	}
}

func (c *Calculator) Pow(n float64) float64 {
	c.addHistory(powOp, []float64{n})
	return c.powOp(n)
}

func (c *Calculator) powOp(n float64) float64 {
	c.current = math.Pow(c.current, n)
	return c.current
}

func (c *Calculator) Cancel() float64 {
	c.resetHistory()
	return c.cancelOp()
}

func (c *Calculator) cancelOp() float64 {
	c.current = 0
	return c.current
}

func (c *Calculator) Repeat(n float64) (float64, error) {
	if n < 0 {
		return c.current, errors.New("can't use negative number")
	}

	res, err := c.repeatFrom(n, len(c.history))
	if err != nil {
		return 0, err
	}

	return res, nil
}

// repeatFrom takes n as how many commands that will be repeated and from denoted that repetition will start from index-1th command
func (c *Calculator) repeatFrom(n float64, from int) (float64, error) {
	rewind := int(n)
	startRepeat := from - rewind
	if startRepeat < 0 {
		startRepeat = 0
	}
	for i := startRepeat; i < from && i < len(c.history); i++ {
		command := c.history[i]
		if command == nil {
			continue
		}

		var err error

		args := command.args
		switch command.op {
		case addOp:
			c.Add(args[0])
		case subtractOp:
			c.Subtract(args[0])
		case multiplyOp:
			c.Multiply(args[0])
		case divideOp:
			c.Divide(args[0])
		case negOp:
			c.Neg()
		case rootOp:
			c.Root(args[0])
		case powOp:
			c.Pow(args[0])
		case absOp:
			c.Abs()
		}

		if err != nil {
			return c.current, err
		}
	}

	return c.current, nil
}

func (c *Calculator) addHistory(op string, args []float64) {
	c.history = append(c.history, &command{op, args})
}

func (c *Calculator) resetHistory() {
	c.history = make([]*command, 0)
}

func (c *Calculator) readHistory() []string {
	hist := make([]string, len(c.history))
	for i, c := range c.history {
		hist[i] = fmt.Sprintf("%s - %v", c.op, c.args)
	}

	return hist
}
