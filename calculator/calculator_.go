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

func (c *Calculator) AddCurrent(a float64) float64 {
	c.addHistory(addOp, []float64{a})
	return c.addOpCurrent(a)
}

func (c *Calculator) addOpCurrent(a float64) float64 {
	c.current = c.current + a
	return c.current
}

func (c *Calculator) SubtractCurrent(a float64) float64 {
	c.addHistory(subtractOp, []float64{a})
	return c.subtractOpCurrent(a)
}

func (c *Calculator) subtractOpCurrent(a float64) float64 {
	c.current = c.current - a
	return c.current
}

func (c *Calculator) MultiplyCurrent(a float64) float64 {
	c.addHistory(multiplyOp, []float64{a})
	return c.multiplyOpCurrent(a)
}

func (c *Calculator) multiplyOpCurrent(a float64) float64 {
	c.current = c.current * a
	return c.current
}

func (c *Calculator) DivideCurrent(a float64) float64 {
	c.addHistory(divideOp, []float64{a})
	return c.divideOpCurrent(a)
}

func (c *Calculator) divideOpCurrent(a float64) float64 {
	if a == 0 {
		c.current = math.NaN()
	} else {
		c.current = c.current / a
	}

	return c.current
}

func (c *Calculator) NegCurrent() float64 {
	c.addHistory(negOp, []float64{})
	return c.negOpCurrent()
}

func (c *Calculator) negOpCurrent() float64 {
	c.current = c.multiplyOpCurrent(-1)
	return c.current
}

func (c *Calculator) AbsCurrent() float64 {
	c.addHistory(absOp, []float64{})
	return c.absOpCurrent()
}

func (c *Calculator) absOpCurrent() float64 {
	c.current = math.Abs(c.current)
	return c.current
}

func (c *Calculator) RootCurrent(n float64) float64 {
	c.addHistory(rootOp, []float64{n})
	return c.rootOpCurrent(n)
}

func (c *Calculator) rootOpCurrent(n float64) float64 {
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

func (c *Calculator) PowCurrent(n float64) float64 {
	c.addHistory(powOp, []float64{n})
	return c.powOpCurrent(n)
}

func (c *Calculator) powOpCurrent(n float64) float64 {
	c.current = math.Pow(c.current, n)
	return c.current
}

func (c *Calculator) Cancel() float64 {
	c.addHistory(cancelOp, []float64{})
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

	// insert history after successfully computed
	c.addHistory(repeatOp, []float64{n})

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
			c.addOpCurrent(args[0])
		case subtractOp:
			c.subtractOpCurrent(args[0])
		case multiplyOp:
			c.multiplyOpCurrent(args[0])
		case divideOp:
			c.divideOpCurrent(args[0])
		case negOp:
			c.negOpCurrent()
		case rootOp:
			c.rootOpCurrent(args[0])
		case powOp:
			c.powOpCurrent(args[0])
		case cancelOp:
			c.cancelOp()
		case absOp:
			c.absOpCurrent()
		case repeatOp:
			_, err = c.repeatFrom(args[0], i) // repeat from itself
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

func (c *Calculator) readHistory() []string {
	hist := make([]string, len(c.history))
	for i, c := range c.history {
		hist[i] = fmt.Sprintf("%s - %v", c.op, c.args)
	}

	return hist
}
