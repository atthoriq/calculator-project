package calculator

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCalculator_Add(t *testing.T) {
	type args struct {
		a float64
	}
	tests := []struct {
		name           string
		args           args
		want           float64
		preExpectation func(c *NewCalculator)
	}{
		{
			name: "input is maxfloat and current is 0 - return maxfloat64",
			args: args{
				a: math.MaxFloat64,
			},
			want:           math.MaxFloat64,
			preExpectation: func(c *NewCalculator) {},
		},
		{
			name: "input is maxfloat and current is maxfloat - return +inf",
			args: args{
				a: math.MaxFloat64,
			},
			want: math.Inf(1),
			preExpectation: func(c *NewCalculator) {
				c.Add(math.MaxFloat64)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InitNewCalculator()
			tt.preExpectation(c)
			if got := c.Add(tt.args.a); !floatEqual(got.GetResult(), tt.want) {
				t.Errorf("Calculator.Add() = %v, want %v", got.GetResult(), tt.want)
			}
		})
	}
}

func TestNewCalculator_SubtractCurrent(t *testing.T) {
	type args struct {
		a float64
	}
	tests := []struct {
		name           string
		args           args
		preExpectation func(c *NewCalculator)
		want           float64
	}{
		{
			name: "input is maxfloat and current is 0 - return -maxfloat64",
			args: args{
				a: math.MaxFloat64,
			},
			want: -math.MaxFloat64,
			preExpectation: func(c *NewCalculator) {
			},
		},
		{
			name: "input is maxfloat and current is -maxfloat - return -inf",
			args: args{
				a: math.MaxFloat64,
			},
			want: math.Inf(-1),
			preExpectation: func(c *NewCalculator) {
				c.Subtract(math.MaxFloat64)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InitNewCalculator()
			tt.preExpectation(c)
			if got := c.Subtract(tt.args.a); !floatEqual(got.GetResult(), tt.want) {
				t.Errorf("Calculator.Subtract() = %v, want %v", got.GetResult(), tt.want)
			}
		})
	}
}

func TestNewCalculator_Multiply(t *testing.T) {
	type args struct {
		a float64
	}
	tests := []struct {
		name           string
		args           args
		want           float64
		preExpectation func(c *NewCalculator)
	}{
		{
			name: "input is maxfloat and current is maxfloat - return +inf",
			args: args{
				a: math.MaxFloat64,
			},
			want: math.Inf(1),
			preExpectation: func(c *NewCalculator) {
				c.Add(math.MaxFloat64)
			},
		},
		{
			name: "input is 0 and current is 0 - return 0",
			args: args{
				a: 0,
			},
			want: 0,
			preExpectation: func(c *NewCalculator) {
			},
		},
		{
			name: "input is -0 and current is non-zero number - return 0",
			args: args{
				a: -0,
			},
			want: 0,
			preExpectation: func(c *NewCalculator) {
				c.Add(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InitNewCalculator()
			tt.preExpectation(c)
			if got := c.Multiply(tt.args.a); !floatEqual(got.GetResult(), tt.want) {
				t.Errorf("Calculator.Multiply() = %v, want %v", got.GetResult(), tt.want)
			}
		})
	}
}

func TestNewCalculator_DivideCurrent(t *testing.T) {
	type args struct {
		a float64
	}
	tests := []struct {
		name           string
		args           args
		want           float64
		preExpectation func(c *NewCalculator)
	}{
		{
			name: "current non-zero number divided by 0 - return NaN",
			args: args{
				a: 0,
			},
			want: math.NaN(),
			preExpectation: func(c *NewCalculator) {
				c.Add(1)
			},
		},
		{
			name: "current non-zero number divided by -0 - return NaN",
			args: args{
				a: -0,
			},
			want: math.NaN(),
			preExpectation: func(c *NewCalculator) {
				c.Add(1)
			},
		},
		{
			name: "current 1 number divided by maxfloat - return 5.562684646268003e-309",
			args: args{
				a: math.MaxFloat64,
			},
			want: 5.562684646268003e-309,
			preExpectation: func(c *NewCalculator) {
				c.Add(1)
			},
		},
		{
			name: "current smallest number float64 number divided by maxfloat - return 0",
			args: args{
				a: math.MaxFloat64,
			},
			want: 0,
			preExpectation: func(c *NewCalculator) {
				c.Add(math.SmallestNonzeroFloat64)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InitNewCalculator()
			tt.preExpectation(c)
			if got := c.Divide(tt.args.a); !floatEqual(got.GetResult(), tt.want) {
				t.Errorf("Calculator.Divide() = %v, want %v", got.GetResult(), tt.want)
			}
		})
	}
}

func TestNewCalculator_AbsCurrent(t *testing.T) {
	tests := []struct {
		name           string
		want           float64
		preExpectation func(c *NewCalculator)
	}{
		{
			name: "current is positive",
			want: math.MaxFloat64,
			preExpectation: func(c *NewCalculator) {
				c.Add(math.MaxFloat64).GetResult()
			},
		},
		{
			name: "current is negative",
			want: math.MaxFloat64,
			preExpectation: func(c *NewCalculator) {
				c.Subtract(math.MaxFloat64).GetResult()
			},
		},
		{
			name:           "current is 0",
			want:           0,
			preExpectation: func(c *NewCalculator) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InitNewCalculator()
			tt.preExpectation(c)
			if got := c.Abs(); !floatEqual(got.GetResult(), tt.want) {
				t.Errorf("Calculator.Abs() = %v, want %v", got.GetResult(), tt.want)
			}
		})
	}
}

func TestNewCalculator_RootCurrent(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name           string
		args           args
		want           float64
		preExpectation func(c *NewCalculator)
	}{
		{
			name: "square root of maxfloat - return a sqrt of maxfloat",
			args: args{
				a: 2,
			},
			want: 1.3407807929942596e+154,
			preExpectation: func(c *NewCalculator) {
				c.Add(math.MaxFloat64)
			},
		},
		{
			name: "square root of 0 - return 0",
			args: args{
				a: 2,
			},
			want:           0,
			preExpectation: func(c *NewCalculator) {},
		},
		{
			name: "cube root of 0 - return 0",
			args: args{
				a: 3,
			},
			want:           0,
			preExpectation: func(c *NewCalculator) {},
		},
		{
			name: "quartic root of 0 - return NaN (not supported)",
			args: args{
				a: 4,
			},
			want:           math.NaN(),
			preExpectation: func(c *NewCalculator) {},
		},
		{
			name: "maxint root of 0 - return NaN (not supported)",
			args: args{
				a: math.MaxInt,
			},
			want:           math.NaN(),
			preExpectation: func(c *NewCalculator) {},
		},
		{
			name: "-maxint root of 0 - return NaN (not supported)",
			args: args{
				a: math.MinInt,
			},
			want:           math.NaN(),
			preExpectation: func(c *NewCalculator) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InitNewCalculator()
			tt.preExpectation(c)
			if got := c.Root(tt.args.a); !floatEqual(got.GetResult(), tt.want) {
				t.Errorf("Calculator.Root() = %v, want %v", got.GetResult(), tt.want)
			}
		})
	}
}

func TestNewCalculator_PowCurrent(t *testing.T) {
	type args struct {
		a float64
	}
	tests := []struct {
		name           string
		args           args
		want           float64
		preExpectation func(c *NewCalculator)
	}{
		{
			name: "0 pow of maxfloat64 - return 0",
			args: args{
				a: math.MaxFloat64,
			},
			want:           0,
			preExpectation: func(c *NewCalculator) {},
		},
		{
			name: "maxfloat64 pow of maxfloat64 - return +inf",
			args: args{
				a: math.MaxFloat64,
			},
			want: math.Inf(1),
			preExpectation: func(c *NewCalculator) {
				c.Add(math.MaxFloat64)
			},
		},
		{
			name: "maxfloat64 pow of -maxfloat64 - return 0",
			args: args{
				a: -math.MaxFloat64,
			},
			want: 0,
			preExpectation: func(c *NewCalculator) {
				c.Add(math.MaxFloat64)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InitNewCalculator()
			tt.preExpectation(c)
			if got := c.Pow(tt.args.a); !floatEqual(got.GetResult(), tt.want) {
				t.Errorf("Calculator.Pow() = %v, want %v", got.GetResult(), tt.want)
			}
		})
	}
}

func TestNewCalculator_Cancel(t *testing.T) {
	tests := []struct {
		name           string
		want           float64
		preExpectation func(c *NewCalculator)
		expectation    func(c *NewCalculator)
	}{
		{
			name: "current 0 - return 0",
			want: 0,
			expectation: func(c *NewCalculator) {
				assert.Len(t, c.history, 0)
			},
		},
		{
			name: "current smallnonzero - return 0",
			want: 0,
			preExpectation: func(c *NewCalculator) {
				c.Add(math.SmallestNonzeroFloat64).GetResult()
			},
			expectation: func(c *NewCalculator) {
				assert.Len(t, c.history, 0)
			},
		},
		{
			name: "current non zero - return 0",
			want: 0,
			preExpectation: func(c *NewCalculator) {
				c.Add(1).GetResult()
			},
			expectation: func(c *NewCalculator) {
				assert.Len(t, c.history, 0)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InitNewCalculator()
			if got := c.Cancel(); !floatEqual(got.GetResult(), tt.want) {
				t.Errorf("Calculator.Cancel() = %v, want %v", got.GetResult(), tt.want)
			}
			tt.expectation(c)
		})
	}
}

func TestNewCalculator_Repeat(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name           string
		args           args
		want           float64
		preExpectation func(c *NewCalculator)
		expectation    func(c *NewCalculator)
	}{
		{
			name: "repeat normal",
			preExpectation: func(c *NewCalculator) {
				c.Add(2).
					Add(5).
					GetResult()
			},
			args: args{
				a: 2,
			},
			want: 14,
			expectation: func(c *NewCalculator) {
				assert.Len(t, c.history, 4)
			},
		},
		{
			name:           "repeat N but there is no history",
			preExpectation: func(c *NewCalculator) {},
			args: args{
				a: 2,
			},
			want: 0,
			expectation: func(c *NewCalculator) {
				assert.Len(t, c.history, 0)
			},
		},
		{
			name: "repeat N but N is bigger than the number of history",
			preExpectation: func(c *NewCalculator) {
				c.Add(2).
					GetResult()
			},
			args: args{
				a: math.MaxInt,
			},
			want: 4,
			expectation: func(c *NewCalculator) {
				assert.Len(t, c.history, 2)
			},
		},
		{
			name: "repeat N but N is negative",
			preExpectation: func(c *NewCalculator) {
				c.Add(2).
					GetResult()
			},
			args: args{
				a: math.MinInt,
			},
			want: 2,
			expectation: func(c *NewCalculator) {
				assert.Len(t, c.history, 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InitNewCalculator()
			tt.preExpectation(c)
			got := c.Repeat(tt.args.a)
			if !floatEqual(got.GetResult(), tt.want) {
				t.Errorf("Calculator.Repeat() = %v, want %v", got.GetResult(), tt.want)
			}
			tt.expectation(c)
		})
	}
}

func TestNewCalculator_Repeat_With_Non_Math_Op(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name                 string
		args                 args
		want                 float64
		wantErr              bool
		preRepeatExpectation func(c *NewCalculator)
	}{
		{
			name: "cancel then repeat - current should be 0",
			args: args{
				a: 2,
			},
			want:    0.0,
			wantErr: false,
			preRepeatExpectation: func(c *NewCalculator) {
				c.Cancel().
					GetResult()
			},
		},
		{
			name: "cancel, doing some mathematical operations, repeat - it will calculate the newest history",
			args: args{
				a: 2,
			},
			want:    2.0,
			wantErr: false,
			preRepeatExpectation: func(c *NewCalculator) {
				c.Cancel().
					Add(1).
					GetResult()
			},
		},
		{
			name: "cancel, doing some mathematical operations, repeat, then repeat - there will be no repeat chain",
			args: args{
				a: 1,
			},
			want:    150.0,
			wantErr: false,
			preRepeatExpectation: func(c *NewCalculator) {
				c.Cancel().
					Add(1).
					Multiply(5).
					Repeat(2).
					GetResult()
			},
		},
		{
			name: "repeat chain",
			args: args{
				a: 3,
			},
			want:    130.0,
			wantErr: false,
			preRepeatExpectation: func(c *NewCalculator) {
				c.Add(5).
					Multiply(2).
					Repeat(2).
					GetResult()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InitNewCalculator()

			tt.preRepeatExpectation(c)
			got := c.Repeat(tt.args.a) // 30 * 5
			if !floatEqual(got.GetResult(), tt.want) {
				t.Errorf("Calculator.Repeat() = %v, want %v", got.GetResult(), tt.want)
			}
		})
	}
}
