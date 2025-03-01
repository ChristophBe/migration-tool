package actions

type ExecutionLogger interface {
	LogExecution([]StepResult) error
	LoadExecutionLog() (ExecutionLogs, error)
}
