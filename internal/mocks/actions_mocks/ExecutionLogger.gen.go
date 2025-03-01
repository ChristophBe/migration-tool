// Code generated by mockery v2.52.3. DO NOT EDIT.

package actions_mocks

import (
	actions "github.com/ChristophBe/migration-tool/pkg/actions"
	mock "github.com/stretchr/testify/mock"
)

// ExecutionLogger is an autogenerated mock type for the ExecutionLogger type
type ExecutionLogger struct {
	mock.Mock
}

type ExecutionLogger_Expecter struct {
	mock *mock.Mock
}

func (_m *ExecutionLogger) EXPECT() *ExecutionLogger_Expecter {
	return &ExecutionLogger_Expecter{mock: &_m.Mock}
}

// LoadExecutionLog provides a mock function with no fields
func (_m *ExecutionLogger) LoadExecutionLog() (actions.ExecutionLogs, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LoadExecutionLog")
	}

	var r0 actions.ExecutionLogs
	var r1 error
	if rf, ok := ret.Get(0).(func() (actions.ExecutionLogs, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() actions.ExecutionLogs); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(actions.ExecutionLogs)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecutionLogger_LoadExecutionLog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadExecutionLog'
type ExecutionLogger_LoadExecutionLog_Call struct {
	*mock.Call
}

// LoadExecutionLog is a helper method to define mock.On call
func (_e *ExecutionLogger_Expecter) LoadExecutionLog() *ExecutionLogger_LoadExecutionLog_Call {
	return &ExecutionLogger_LoadExecutionLog_Call{Call: _e.mock.On("LoadExecutionLog")}
}

func (_c *ExecutionLogger_LoadExecutionLog_Call) Run(run func()) *ExecutionLogger_LoadExecutionLog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ExecutionLogger_LoadExecutionLog_Call) Return(_a0 actions.ExecutionLogs, _a1 error) *ExecutionLogger_LoadExecutionLog_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ExecutionLogger_LoadExecutionLog_Call) RunAndReturn(run func() (actions.ExecutionLogs, error)) *ExecutionLogger_LoadExecutionLog_Call {
	_c.Call.Return(run)
	return _c
}

// LogExecution provides a mock function with given fields: _a0
func (_m *ExecutionLogger) LogExecution(_a0 []actions.StepResult) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for LogExecution")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]actions.StepResult) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExecutionLogger_LogExecution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LogExecution'
type ExecutionLogger_LogExecution_Call struct {
	*mock.Call
}

// LogExecution is a helper method to define mock.On call
//   - _a0 []actions.StepResult
func (_e *ExecutionLogger_Expecter) LogExecution(_a0 interface{}) *ExecutionLogger_LogExecution_Call {
	return &ExecutionLogger_LogExecution_Call{Call: _e.mock.On("LogExecution", _a0)}
}

func (_c *ExecutionLogger_LogExecution_Call) Run(run func(_a0 []actions.StepResult)) *ExecutionLogger_LogExecution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]actions.StepResult))
	})
	return _c
}

func (_c *ExecutionLogger_LogExecution_Call) Return(_a0 error) *ExecutionLogger_LogExecution_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ExecutionLogger_LogExecution_Call) RunAndReturn(run func([]actions.StepResult) error) *ExecutionLogger_LogExecution_Call {
	_c.Call.Return(run)
	return _c
}

// NewExecutionLogger creates a new instance of ExecutionLogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExecutionLogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExecutionLogger {
	mock := &ExecutionLogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
