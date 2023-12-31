// Code generated by MockGen. DO NOT EDIT.
// Source: calculator/new_calculator.go

// Package mock_calculator is a generated GoMock package.
package mock_calculator

import (
	gomock "github.com/golang/mock/gomock"
	calculator "gitlab.com/atthoriq/calculator-project/calculator"
	reflect "reflect"
)

// MockNewCalculator is a mock of NewCalculator interface
type MockNewCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockNewCalculatorMockRecorder
}

// MockNewCalculatorMockRecorder is the mock recorder for MockNewCalculator
type MockNewCalculatorMockRecorder struct {
	mock *MockNewCalculator
}

// NewMockNewCalculator creates a new mock instance
func NewMockNewCalculator(ctrl *gomock.Controller) *MockNewCalculator {
	mock := &MockNewCalculator{ctrl: ctrl}
	mock.recorder = &MockNewCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNewCalculator) EXPECT() *MockNewCalculatorMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockNewCalculator) Add(a float64) calculator.NewCalculator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", a)
	ret0, _ := ret[0].(calculator.NewCalculator)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockNewCalculatorMockRecorder) Add(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockNewCalculator)(nil).Add), a)
}

// Subtract mocks base method
func (m *MockNewCalculator) Subtract(a float64) calculator.NewCalculator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subtract", a)
	ret0, _ := ret[0].(calculator.NewCalculator)
	return ret0
}

// Subtract indicates an expected call of Subtract
func (mr *MockNewCalculatorMockRecorder) Subtract(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subtract", reflect.TypeOf((*MockNewCalculator)(nil).Subtract), a)
}

// Multiply mocks base method
func (m *MockNewCalculator) Multiply(a float64) calculator.NewCalculator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Multiply", a)
	ret0, _ := ret[0].(calculator.NewCalculator)
	return ret0
}

// Multiply indicates an expected call of Multiply
func (mr *MockNewCalculatorMockRecorder) Multiply(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Multiply", reflect.TypeOf((*MockNewCalculator)(nil).Multiply), a)
}

// Divide mocks base method
func (m *MockNewCalculator) Divide(a float64) calculator.NewCalculator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Divide", a)
	ret0, _ := ret[0].(calculator.NewCalculator)
	return ret0
}

// Divide indicates an expected call of Divide
func (mr *MockNewCalculatorMockRecorder) Divide(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Divide", reflect.TypeOf((*MockNewCalculator)(nil).Divide), a)
}

// Abs mocks base method
func (m *MockNewCalculator) Abs() calculator.NewCalculator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Abs")
	ret0, _ := ret[0].(calculator.NewCalculator)
	return ret0
}

// Abs indicates an expected call of Abs
func (mr *MockNewCalculatorMockRecorder) Abs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abs", reflect.TypeOf((*MockNewCalculator)(nil).Abs))
}

// Root mocks base method
func (m *MockNewCalculator) Root(a int) calculator.NewCalculator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Root", a)
	ret0, _ := ret[0].(calculator.NewCalculator)
	return ret0
}

// Root indicates an expected call of Root
func (mr *MockNewCalculatorMockRecorder) Root(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Root", reflect.TypeOf((*MockNewCalculator)(nil).Root), a)
}

// Pow mocks base method
func (m *MockNewCalculator) Pow(a float64) calculator.NewCalculator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pow", a)
	ret0, _ := ret[0].(calculator.NewCalculator)
	return ret0
}

// Pow indicates an expected call of Pow
func (mr *MockNewCalculatorMockRecorder) Pow(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pow", reflect.TypeOf((*MockNewCalculator)(nil).Pow), a)
}

// Repeat mocks base method
func (m *MockNewCalculator) Repeat(a int) calculator.NewCalculator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Repeat", a)
	ret0, _ := ret[0].(calculator.NewCalculator)
	return ret0
}

// Repeat indicates an expected call of Repeat
func (mr *MockNewCalculatorMockRecorder) Repeat(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Repeat", reflect.TypeOf((*MockNewCalculator)(nil).Repeat), a)
}

// Cancel mocks base method
func (m *MockNewCalculator) Cancel() calculator.NewCalculator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel")
	ret0, _ := ret[0].(calculator.NewCalculator)
	return ret0
}

// Cancel indicates an expected call of Cancel
func (mr *MockNewCalculatorMockRecorder) Cancel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockNewCalculator)(nil).Cancel))
}

// GetResult mocks base method
func (m *MockNewCalculator) GetResult() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResult")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetResult indicates an expected call of GetResult
func (mr *MockNewCalculatorMockRecorder) GetResult() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResult", reflect.TypeOf((*MockNewCalculator)(nil).GetResult))
}
