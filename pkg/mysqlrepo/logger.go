package mysqlrepo

import "log"

const (
	LogFormat = "Query: %s, Args: %v"
)

type Loggerer interface {
	Printf(stmt string, args ...any)
	Error(v ...any)
}

type defaultLogger struct{}

func (l defaultLogger) Printf(stmt string, args ...any) {
	log.Printf(LogFormat, stmt, args)
}

func (l defaultLogger) Error(v ...any) {
	log.Println(v...)
}

func DefaultLogger() defaultLogger {
	return defaultLogger{}
}
