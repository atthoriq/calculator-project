package calculator

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func floatEqual(a, b float64) bool {
	if math.IsNaN(a) && math.IsNaN(b) {
		return true
	}
	return a == b
}

func NewCalculator_Add(t *testing.T) {
	type fields struct {
		history []*command
		current float64
	}
	type args struct {
		a float64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        float64
		expectation func(c *Calculator)
	}{
		{
			name: "input is maxfloat and current is 0 - return maxfloat64",
			fields: fields{
				history: []*command{},
				current: 0,
			},
			args: args{
				a: math.MaxFloat64,
			},
			want: math.MaxFloat64,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
			},
		},
		{
			name: "input is maxfloat and current is maxfloat - return +inf",
			fields: fields{
				history: []*command{},
				current: math.MaxFloat64,
			},
			args: args{
				a: math.MaxFloat64,
			},
			want: math.Inf(1),
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, addOp)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				history: tt.fields.history,
				current: tt.fields.current,
			}
			if got := c.Add(tt.args.a); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.Add() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func NewCalculator_Subtract(t *testing.T) {
	type fields struct {
		history []*command
		current float64
	}
	type args struct {
		a float64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        float64
		expectation func(c *Calculator)
	}{
		{
			name: "input is maxfloat and current is 0 - return -maxfloat64",
			fields: fields{
				history: []*command{},
				current: 0,
			},
			args: args{
				a: math.MaxFloat64,
			},
			want: -math.MaxFloat64,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, subtractOp)
			},
		},
		{
			name: "input is maxfloat and current is -maxfloat - return -inf",
			fields: fields{
				history: []*command{},
				current: -math.MaxFloat64,
			},
			args: args{
				a: math.MaxFloat64,
			},
			want: math.Inf(-1),
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, subtractOp)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				history: tt.fields.history,
				current: tt.fields.current,
			}
			if got := c.Subtract(tt.args.a); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.Subtract() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func NewCalculator_Multiply(t *testing.T) {
	type fields struct {
		history []*command
		current float64
	}
	type args struct {
		a float64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        float64
		expectation func(c *Calculator)
	}{
		{
			name: "input is maxfloat and current is maxfloat - return +inf",
			fields: fields{
				history: []*command{},
				current: math.MaxFloat64,
			},
			args: args{
				a: math.MaxFloat64,
			},
			want: math.Inf(1),
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, multiplyOp)
			},
		},
		{
			name: "input is 0 and current is 0 - return 0",
			fields: fields{
				history: []*command{},
				current: 0,
			},
			args: args{
				a: 0,
			},
			want: 0,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, multiplyOp)
			},
		},
		{
			name: "input is -0 and current is non-zero number - return 0",
			fields: fields{
				history: []*command{},
				current: 5,
			},
			args: args{
				a: -0,
			},
			want: 0,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, multiplyOp)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				history: tt.fields.history,
				current: tt.fields.current,
			}
			if got := c.Multiply(tt.args.a); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.Multiply() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func NewCalculator_Divide(t *testing.T) {
	type fields struct {
		history []*command
		current float64
	}
	type args struct {
		a float64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        float64
		expectation func(c *Calculator)
	}{
		{
			name: "current non-zero number divided by 0 - return NaN",
			fields: fields{
				history: []*command{},
				current: math.MaxFloat64,
			},
			args: args{
				a: 0,
			},
			want: math.NaN(),
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, divideOp)
			},
		},
		{
			name: "current non-zero number divided by -0 - return NaN",
			fields: fields{
				history: []*command{},
				current: math.MaxFloat64,
			},
			args: args{
				a: -0,
			},
			want: math.NaN(),
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, divideOp)
			},
		},
		{
			name: "current 1 number divided by maxfloat - return 5.562684646268003e-309",
			fields: fields{
				history: []*command{},
				current: 1,
			},
			args: args{
				a: math.MaxFloat64,
			},
			want: 5.562684646268003e-309,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, divideOp)
			},
		},
		{
			name: "current smallest number float64 number divided by maxfloat - return 0",
			fields: fields{
				history: []*command{},
				current: math.SmallestNonzeroFloat64,
			},
			args: args{
				a: math.MaxFloat64,
			},
			want: 0,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, divideOp)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				history: tt.fields.history,
				current: tt.fields.current,
			}
			if got := c.Divide(tt.args.a); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.Divide() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func NewCalculator_Neg(t *testing.T) {
	type fields struct {
		history []*command
		current float64
	}
	tests := []struct {
		name        string
		fields      fields
		want        float64
		expectation func(c *Calculator)
	}{
		{
			name: "current is positive",
			fields: fields{
				history: []*command{},
				current: math.MaxFloat64,
			},
			want: -math.MaxFloat64,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, negOp)
			},
		},
		{
			name: "current is negative",
			fields: fields{
				history: []*command{},
				current: -math.MaxFloat64,
			},
			want: math.MaxFloat64,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, negOp)
			},
		},
		{
			name: "current is 0",
			fields: fields{
				history: []*command{},
				current: 0,
			},
			want: 0,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, negOp)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				history: tt.fields.history,
				current: tt.fields.current,
			}
			if got := c.Neg(); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.Neg() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func NewCalculator_Abs(t *testing.T) {
	type fields struct {
		history []*command
		current float64
	}
	tests := []struct {
		name        string
		fields      fields
		want        float64
		expectation func(c *Calculator)
	}{
		{
			name: "current is positive",
			fields: fields{
				history: []*command{},
				current: math.MaxFloat64,
			},
			want: math.MaxFloat64,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, absOp)
			},
		},
		{
			name: "current is negative",
			fields: fields{
				history: []*command{},
				current: -math.MaxFloat64,
			},
			want: math.MaxFloat64,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, absOp)
			},
		},
		{
			name: "current is 0",
			fields: fields{
				history: []*command{},
				current: 0,
			},
			want: 0,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, absOp)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				history: tt.fields.history,
				current: tt.fields.current,
			}
			if got := c.Abs(); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.Abs() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func NewCalculator_Root(t *testing.T) {
	type fields struct {
		history []*command
		current float64
	}
	type args struct {
		a float64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        float64
		expectation func(c *Calculator)
	}{
		{
			name: "square root of maxfloat - return a sqrt of maxfloat",
			fields: fields{
				history: []*command{},
				current: math.MaxFloat64,
			},
			args: args{
				a: 2,
			},
			want: 1.3407807929942596e+154,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, rootOp)
			},
		},
		{
			name: "square root of 0 - return 0",
			fields: fields{
				history: []*command{},
				current: 0,
			},
			args: args{
				a: 2,
			},
			want: 0,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, rootOp)
			},
		},
		{
			name: "cube root of 0 - return 0",
			fields: fields{
				history: []*command{},
				current: 0,
			},
			args: args{
				a: 3,
			},
			want: 0,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, rootOp)
			},
		},
		{
			name: "quartic root of 0 - return NaN (not supported)",
			fields: fields{
				history: []*command{},
				current: 0,
			},
			args: args{
				a: 4,
			},
			want: math.NaN(),
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, rootOp)
			},
		},
		{
			name: "maxfloat root of 0 - return NaN (not supported)",
			fields: fields{
				history: []*command{},
				current: 0,
			},
			args: args{
				a: math.MaxFloat64,
			},
			want: math.NaN(),
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, rootOp)
			},
		},
		{
			name: "-maxfloat root of 0 - return NaN (not supported)",
			fields: fields{
				history: []*command{},
				current: 0,
			},
			args: args{
				a: -math.MaxFloat64,
			},
			want: math.NaN(),
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, rootOp)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				history: tt.fields.history,
				current: tt.fields.current,
			}
			if got := c.Root(tt.args.a); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.Root() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func NewCalculator_Pow(t *testing.T) {
	type fields struct {
		history []*command
		current float64
	}
	type args struct {
		a float64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        float64
		expectation func(c *Calculator)
	}{
		{
			name: "0 pow of maxfloat64 - return 0",
			fields: fields{
				history: []*command{},
				current: 0,
			},
			args: args{
				a: math.MaxFloat64,
			},
			want: 0,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, powOp)
			},
		},
		{
			name: "maxfloat64 pow of maxfloat64 - return +inf",
			fields: fields{
				history: []*command{},
				current: math.MaxFloat64,
			},
			args: args{
				a: math.MaxFloat64,
			},
			want: math.Inf(1),
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, powOp)
			},
		},
		{
			name: "maxfloat64 pow of -maxfloat64 - return 0",
			fields: fields{
				history: []*command{},
				current: math.MaxFloat64,
			},
			args: args{
				a: -math.MaxFloat64,
			},
			want: 0,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, powOp)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				history: tt.fields.history,
				current: tt.fields.current,
			}
			if got := c.Pow(tt.args.a); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.Pow() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func NewCalculator_Cancel(t *testing.T) {
	type fields struct {
		history []*command
		current float64
	}
	tests := []struct {
		name        string
		fields      fields
		want        float64
		expectation func(c *Calculator)
	}{
		{
			name: "current 0 - return 0",
			fields: fields{
				history: []*command{},
				current: 0,
			},
			want: 0,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 0)
			},
		},
		{
			name: "current smallnonzero - return 0",
			fields: fields{
				history: []*command{},
				current: math.SmallestNonzeroFloat64,
			},
			want: 0,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 0)
			},
		},
		{
			name: "current non zero - return 0",
			fields: fields{
				history: []*command{
					{
						op: addOp,
						args: []float64{
							1,
						},
					},
				},
				current: 1,
			},
			want: 0,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 0)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				history: tt.fields.history,
				current: tt.fields.current,
			}
			if got := c.Cancel(); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.Cancel() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func NewCalculator_Repeat(t *testing.T) {
	type fields struct {
		history []*command
		current float64
	}
	type args struct {
		a float64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        float64
		wantErr     bool
		expectation func(c *Calculator)
	}{
		{
			name: "repeat normal",
			fields: fields{
				history: []*command{
					{
						op:   addOp,
						args: []float64{5},
					},
					{
						op:   addOp,
						args: []float64{2},
					},
				},
				current: 7,
			},
			args: args{
				a: 2,
			},
			want:    14,
			wantErr: false,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 4)
			},
		},
		{
			name: "repeat N but there is no history",
			fields: fields{
				history: []*command{},
				current: 0,
			},
			args: args{
				a: 2,
			},
			want:    0,
			wantErr: false,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 0)
			},
		},
		{
			name: "repeat N but N is bigger than the number of history",
			fields: fields{
				history: []*command{
					{
						op:   addOp,
						args: []float64{2},
					},
				},
				current: 2,
			},
			args: args{
				a: math.MaxFloat64,
			},
			want:    4,
			wantErr: false,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 2)
			},
		},
		{
			name: "repeat N but N is negative",
			fields: fields{
				history: []*command{
					{
						op:   addOp,
						args: []float64{2},
					},
				},
				current: 2,
			},
			args: args{
				a: -math.MaxFloat64,
			},
			want:    2,
			wantErr: true,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				history: tt.fields.history,
				current: tt.fields.current,
			}
			got, err := c.Repeat(tt.args.a)
			if !floatEqual(got, tt.want) {
				t.Errorf("Calculator.Repeat() = %v, want %v", got, tt.want)
			}
			if tt.wantErr != (err != nil) {
				t.Errorf("Calculator.Repeat() err = %v", err)
			}
			tt.expectation(c)
		})
	}
}

func NewCalculator_Repeat_With_Non_Math_Op(t *testing.T) {
	type fields struct {
		history []*command
		current float64
	}
	type args struct {
		a float64
	}
	tests := []struct {
		name                 string
		fields               fields
		args                 args
		want                 float64
		wantErr              bool
		preRepeatExpectation func(c *Calculator)
	}{
		{
			name: "cancel then repeat - current should be 0",
			fields: fields{
				history: []*command{
					{
						op:   addOp,
						args: []float64{5},
					},
					{
						op:   addOp,
						args: []float64{2},
					},
				},
				current: 7,
			},
			args: args{
				a: 2,
			},
			want:    0.0,
			wantErr: false,
			preRepeatExpectation: func(c *Calculator) {
				c.Cancel()
			},
		},
		{
			name: "cancel, doing some mathematical operations, repeat - it will calculate the newest history",
			fields: fields{
				history: []*command{
					{
						op:   addOp,
						args: []float64{5},
					},
					{
						op:   addOp,
						args: []float64{2},
					},
				},
				current: 7,
			},
			args: args{
				a: 2,
			},
			want:    2.0,
			wantErr: false,
			preRepeatExpectation: func(c *Calculator) {
				c.Cancel()
				c.Add(1)
			},
		},
		{
			name: "cancel, doing some mathematical operations, repeat, then repeat - there will be no repeat chain",
			fields: fields{
				history: []*command{
					{
						op:   addOp,
						args: []float64{5},
					},
					{
						op:   addOp,
						args: []float64{2},
					},
				},
				current: 7,
			},
			args: args{
				a: 1,
			},
			want:    150.0,
			wantErr: false,
			preRepeatExpectation: func(c *Calculator) {
				c.Cancel()
				c.Add(1)
				c.Multiply(5) // 5
				c.Repeat(2)   // 30 -> add 1 multiply 5 add 1 multiply 5
				// repeat 3 -> add 1 multiply 5 add 1 multiply 5 + multiply 5 add 1 multiply 5
			},
		},
		{
			name: "repeat chain",
			fields: fields{
				history: []*command{
					{
						op:   addOp,
						args: []float64{5},
					},
					{
						op:   multiplyOp,
						args: []float64{2},
					},
				},
				current: 10,
			},
			args: args{
				a: 3,
			},
			want:    130.0,
			wantErr: false,
			preRepeatExpectation: func(c *Calculator) {
				c.Repeat(2) // 10 + 5 * 2 = 30 -> add 5; multuply 2; add 5; multiply 2
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				history: tt.fields.history,
				current: tt.fields.current,
			}

			tt.preRepeatExpectation(c)
			got, err := c.Repeat(tt.args.a) // 30 * 5
			if !floatEqual(got, tt.want) {
				t.Errorf("Calculator.Repeat() = %v, want %v", got, tt.want)
			}
			if tt.wantErr != (err != nil) {
				t.Errorf("Calculator.Repeat() err = %v", err)
			}
		})
	}
}
