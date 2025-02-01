package systemder

type Logger interface {
	Error(format string, args ...any)
}

type NullLogger struct{}

func (l *NullLogger) Error(format string, args ...any) {}
