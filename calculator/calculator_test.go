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

func TestCalculator_AddCurrent(t *testing.T) {
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
			if got := c.AddCurrent(tt.args.a); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.AddCurrent() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func TestCalculator_SubtractCurrent(t *testing.T) {
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
			if got := c.SubtractCurrent(tt.args.a); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.SubtractCurrent() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func TestCalculator_MultiplyCurrent(t *testing.T) {
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
			if got := c.MultiplyCurrent(tt.args.a); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.MultiplyCurrent() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func TestCalculator_DivideCurrent(t *testing.T) {
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
			if got := c.DivideCurrent(tt.args.a); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.DivideCurrent() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func TestCalculator_NegCurrent(t *testing.T) {
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
			if got := c.NegCurrent(); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.NegCurrent() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func TestCalculator_AbsCurrent(t *testing.T) {
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
			if got := c.AbsCurrent(); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.AbsCurrent() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func TestCalculator_RootCurrent(t *testing.T) {
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
			if got := c.RootCurrent(tt.args.a); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.RootCurrent() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func TestCalculator_PowCurrent(t *testing.T) {
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
			if got := c.PowCurrent(tt.args.a); !floatEqual(got, tt.want) {
				t.Errorf("Calculator.PowCurrent() = %v, want %v", got, tt.want)
			}
			tt.expectation(c)
		})
	}
}

func TestCalculator_Cancel(t *testing.T) {
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
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, cancelOp)
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
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, cancelOp)
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

func TestCalculator_Repeat(t *testing.T) {
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
				assert.Len(t, c.history, 3)
				assert.Equal(t, c.history[2].op, repeatOp)
			},
		},
		{
			name: "repeat normal with cancel",
			fields: fields{
				history: []*command{
					{
						op:   addOp,
						args: []float64{5},
					},
					{
						op:   cancelOp,
						args: []float64{},
					},
					{
						op:   addOp,
						args: []float64{2},
					},
				},
				current: 2,
			},
			args: args{
				a: 3,
			},
			want:    2,
			wantErr: false,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 4)
				assert.Equal(t, c.history[3].op, repeatOp)
			},
		},
		{
			name: "repeat the repeat command",
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
					{
						op:   repeatOp,
						args: []float64{2},
					},
				},
				current: 14,
			},
			args: args{
				a: 1,
			},
			want:    21,
			wantErr: false,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 4)
				assert.Equal(t, c.history[3].op, repeatOp)
			},
		},
		{
			name: "repeat N steps but N+m, where m is between N and latest step, there is a repeat command",
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
					{
						op:   repeatOp,
						args: []float64{2},
					},
				},
				current: 14,
			},
			args: args{
				a: 2,
			},
			want:    23,
			wantErr: false,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 4)
				assert.Equal(t, c.history[3].op, repeatOp)
			},
		},
		{
			name: "repeat N steps but N+m, where m is between N and latest step, there is a couple of repeat command",
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
					{
						op:   repeatOp,
						args: []float64{2},
					},
					{
						op:   repeatOp,
						args: []float64{1},
					},
				},
				current: 21,
			},
			args: args{
				a: 2,
			},
			want:    35,
			wantErr: false,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 5)
				assert.Equal(t, c.history[4].op, repeatOp)
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
				assert.Len(t, c.history, 1)
				assert.Equal(t, c.history[0].op, repeatOp)
			},
		},
		{
			name: "repeat N but there is no computation yet, just repeats",
			fields: fields{
				history: []*command{
					{
						op:   repeatOp,
						args: []float64{2},
					},
				},
				current: 0,
			},
			args: args{
				a: 2,
			},
			want:    0,
			wantErr: false,
			expectation: func(c *Calculator) {
				assert.Len(t, c.history, 2)
				assert.Equal(t, c.history[1].op, repeatOp)
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
				assert.Equal(t, c.history[1].op, repeatOp)
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
