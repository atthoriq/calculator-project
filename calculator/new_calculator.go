package calculator

// using pattern builder in functional way
// there's an additional step to return the current result, it is "GetResult" function. This approach results in "currentoperations" variable to hold the calculation before result is returned
// history/currentoperations can be placed in an array of function that handles the calculator struct. each function operation can be different
// history can be written in this package or outside of this package using similar approach.

import (
	"math"
)

type NewCalculator struct {
	current           float64
	currentOperations []operation
	history           []operation
}

type operation func(*NewCalculator)

func InitNewCalculator() *NewCalculator {
	return &NewCalculator{0, []operation{}, []operation{}}
}

func (c *NewCalculator) Add(a float64) *NewCalculator {
	c.currentOperations = append(c.currentOperations, func(nc *NewCalculator) {
		nc.current += a
	})
	return c
}

func (c *NewCalculator) Subtract(a float64) *NewCalculator {
	c.currentOperations = append(c.currentOperations, func(nc *NewCalculator) {
		nc.current -= a
	})
	return c
}

func (c *NewCalculator) Multiply(a float64) *NewCalculator {
	c.currentOperations = append(c.currentOperations, func(nc *NewCalculator) {
		nc.current *= a
	})
	return c
}

func (c *NewCalculator) Divide(a float64) *NewCalculator {
	c.currentOperations = append(c.currentOperations, func(nc *NewCalculator) {
		if a == 0 {
			nc.current = math.NaN()
		} else {
			nc.current /= a
		}
	})
	return c
}

func (c *NewCalculator) Abs() *NewCalculator {
	c.currentOperations = append(c.currentOperations, func(nc *NewCalculator) {
		c.current = math.Abs(c.current)
	})
	return c
}

func (c *NewCalculator) Root(n int) *NewCalculator {
	c.currentOperations = append(c.currentOperations, func(nc *NewCalculator) {
		switch n {
		case 2:
			c.current = math.Sqrt(c.current)
		case 3:
			c.current = math.Cbrt(c.current)
		default:
			c.current = math.NaN()
		}
	})
	return c
}

func (c *NewCalculator) Pow(n float64) *NewCalculator {
	c.currentOperations = append(c.currentOperations, func(nc *NewCalculator) {
		c.current = math.Pow(c.current, n)
	})
	return c
}

func (c *NewCalculator) Cancel() *NewCalculator {
	c.current = 0
	c.currentOperations = []operation{}
	c.history = []operation{}
	return c
}

func (c *NewCalculator) Repeat(n int) *NewCalculator {
	// clean hold operations
	c.GetResult()

	if n < 0 || len(c.history) == 0 {
		return c
	}

	startRepeat := len(c.history) - n
	if startRepeat < 0 {
		startRepeat = 0
	}

	lastNhistory := c.history[startRepeat:]
	for _, op := range lastNhistory {
		op(c)
		c.history = append(c.history, op)
	}
	return c
}

func (c *NewCalculator) GetResult() float64 {
	for _, op := range c.currentOperations {
		op(c)
		c.history = append(c.history, op)
	}
	c.currentOperations = []operation{}
	return c.current
}
