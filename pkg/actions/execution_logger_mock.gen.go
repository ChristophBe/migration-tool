// Code generated by mockery. DO NOT EDIT.

package actions

import mock "github.com/stretchr/testify/mock"

// ExecutionLoggerMock is an autogenerated mock type for the ExecutionLogger type
type ExecutionLoggerMock struct {
	mock.Mock
}

type ExecutionLoggerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ExecutionLoggerMock) EXPECT() *ExecutionLoggerMock_Expecter {
	return &ExecutionLoggerMock_Expecter{mock: &_m.Mock}
}

// LoadExecutionLog provides a mock function with no fields
func (_m *ExecutionLoggerMock) LoadExecutionLog() (ExecutionLogs, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LoadExecutionLog")
	}

	var r0 ExecutionLogs
	var r1 error
	if rf, ok := ret.Get(0).(func() (ExecutionLogs, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() ExecutionLogs); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ExecutionLogs)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecutionLoggerMock_LoadExecutionLog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadExecutionLog'
type ExecutionLoggerMock_LoadExecutionLog_Call struct {
	*mock.Call
}

// LoadExecutionLog is a helper method to define mock.On call
func (_e *ExecutionLoggerMock_Expecter) LoadExecutionLog() *ExecutionLoggerMock_LoadExecutionLog_Call {
	return &ExecutionLoggerMock_LoadExecutionLog_Call{Call: _e.mock.On("LoadExecutionLog")}
}

func (_c *ExecutionLoggerMock_LoadExecutionLog_Call) Run(run func()) *ExecutionLoggerMock_LoadExecutionLog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ExecutionLoggerMock_LoadExecutionLog_Call) Return(_a0 ExecutionLogs, _a1 error) *ExecutionLoggerMock_LoadExecutionLog_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ExecutionLoggerMock_LoadExecutionLog_Call) RunAndReturn(run func() (ExecutionLogs, error)) *ExecutionLoggerMock_LoadExecutionLog_Call {
	_c.Call.Return(run)
	return _c
}

// LogExecution provides a mock function with given fields: _a0
func (_m *ExecutionLoggerMock) LogExecution(_a0 []StepResult) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for LogExecution")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]StepResult) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExecutionLoggerMock_LogExecution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LogExecution'
type ExecutionLoggerMock_LogExecution_Call struct {
	*mock.Call
}

// LogExecution is a helper method to define mock.On call
//   - _a0 []StepResult
func (_e *ExecutionLoggerMock_Expecter) LogExecution(_a0 interface{}) *ExecutionLoggerMock_LogExecution_Call {
	return &ExecutionLoggerMock_LogExecution_Call{Call: _e.mock.On("LogExecution", _a0)}
}

func (_c *ExecutionLoggerMock_LogExecution_Call) Run(run func(_a0 []StepResult)) *ExecutionLoggerMock_LogExecution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]StepResult))
	})
	return _c
}

func (_c *ExecutionLoggerMock_LogExecution_Call) Return(_a0 error) *ExecutionLoggerMock_LogExecution_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ExecutionLoggerMock_LogExecution_Call) RunAndReturn(run func([]StepResult) error) *ExecutionLoggerMock_LogExecution_Call {
	_c.Call.Return(run)
	return _c
}

// NewExecutionLoggerMock creates a new instance of ExecutionLoggerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExecutionLoggerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExecutionLoggerMock {
	mock := &ExecutionLoggerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
