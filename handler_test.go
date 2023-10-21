package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_main "gitlab.com/atthoriq/calculator-project/mock"
)

func Test_calculatorHandler_Handle_Negative_Cases(t *testing.T) {
	ctrl := gomock.NewController(t)
	type fields struct {
		calculator Calculator
	}
	type args struct {
		command string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "commands are empty",
			fields: fields{
				calculator: mock_main.NewMockCalculator(ctrl),
			},
			args: args{
				command: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "commands are more than 2",
			fields: fields{
				calculator: mock_main.NewMockCalculator(ctrl),
			},
			args: args{
				command: "add 20 5",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "commands are 2 but the second argument is not a number",
			fields: fields{
				calculator: mock_main.NewMockCalculator(ctrl),
			},
			args: args{
				command: "add abc",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "commands are 2 but the second argument is a number using comma",
			fields: fields{
				calculator: mock_main.NewMockCalculator(ctrl),
			},
			args: args{
				command: "add 1,2",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "commands are not defined",
			fields: fields{
				calculator: mock_main.NewMockCalculator(ctrl),
			},
			args: args{
				command: "factorial",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "command requires 1 arg but given 2",
			fields: fields{
				calculator: mock_main.NewMockCalculator(ctrl),
			},
			args: args{
				command: "abs 2",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := &calculatorHandler{
				calculator: tt.fields.calculator,
			}
			got, err := ch.Handle(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculatorHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculatorHandler.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

// most test cases will be trivial so they will be ignored as this is written
func Test_calculatorHandler_Handle_Command_Cases(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCalc := mock_main.NewMockCalculator(ctrl)
	type fields struct {
		calculator Calculator
	}
	type args struct {
		command string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        string
		wantErr     bool
		expectation func()
	}{
		{
			name: "commands are with additional spaces",
			fields: fields{
				calculator: mockCalc,
			},
			args: args{
				command: " add  2",
			},
			want:    "2.00",
			wantErr: false,
			expectation: func() {
				mockCalc.EXPECT().Add(float64(2)).Return(float64(2))
			},
		},
		{
			name: "subtract command",
			fields: fields{
				calculator: mockCalc,
			},
			args: args{
				command: "subtract 2",
			},
			want:    "2.00",
			wantErr: false,
			expectation: func() {
				mockCalc.EXPECT().Subtract(float64(2)).Return(float64(2))
			},
		},
		{
			name: "exit command",
			fields: fields{
				calculator: mockCalc,
			},
			args: args{
				command: "exit",
			},
			want:        "",
			wantErr:     false,
			expectation: func() {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := &calculatorHandler{
				calculator: tt.fields.calculator,
			}
			tt.expectation()
			got, err := ch.Handle(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculatorHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculatorHandler.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}
