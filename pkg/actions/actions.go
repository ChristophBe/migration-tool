package actions

type Actions struct {
	executionLogger ExecutionLogger
}

func New(ExecutionLogger ExecutionLogger) *Actions {
	return &Actions{
		executionLogger: ExecutionLogger,
	}
}
