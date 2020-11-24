package stand

import (
	"fmt"
	"time"
)

// Logger writes logs or formatted logs to some kind of output.
type Logger interface {
	LogWriter
	Loggerf
}

// LogWriter writes log strings to some kind of output.
type LogWriter interface {
	// Debug logs a debug message.
	Debug(msg string)
	// Info logs an informational message.
	Info(msg string)
	// Warn logs a warning message.
	Warn(msg string)
	// Error logs an error message.
	Error(msg string)
}

// Loggerf writes formatted logs to some kind of output.
type Loggerf interface {
	// Debugf logs a debug message.
	Debugf(format string, args ...interface{})
	// Infof logs an informational message.
	Infof(format string, args ...interface{})
	// Warnf logs a warning message.
	Warnf(format string, args ...interface{})
	// Errorf logs an error message.
	Errorf(format string, args ...interface{})
}

// NewStdoutLogger creates a Logger that writes to stdout.
func NewStdoutLogger() Logger {
	return NewLogger(stdoutLogWriter{})
}

var _ LogWriter = (*stdoutLogWriter)(nil)

// stdoutLogWriter is a LogWriter that writes to stdout.
type stdoutLogWriter struct{}

// Debug logs a debug message.
func (s stdoutLogWriter) Debug(msg string) {
	s.write("DEBUG", msg)
}

// Info logs an informational message.
func (s stdoutLogWriter) Info(msg string) {
	s.write("INFO", msg)
}

// Warn logs a warning message.
func (s stdoutLogWriter) Warn(msg string) {
	s.write("WARN", msg)
}

// Error logs an error message.
func (s stdoutLogWriter) Error(msg string) {
	s.write("ERROR", msg)
}

// write logs a log line to stdout prefixed by the time and log level.
func (s stdoutLogWriter) write(level string, msg string) {
	fmt.Printf("%s %s: %s\n", time.Now(), level, msg)
}

// NewMultiplexedLogWriter creates a Logger that logs to multiple outputs.
func NewMultiplexedLogWriter(logWriters ...LogWriter) Logger {
	switch {
	case len(logWriters) == 0:
		return noopLogger{}
	case len(logWriters) == 1:
		return NewLogger(logWriters[0])
	default:
		return NewLogger(multiplexedLogWriter{logWriters})
	}
}

// multiplexedLogWriter is a LogWriter that logs to multiple outputs.
type multiplexedLogWriter struct {
	logWriters []LogWriter
}

var _ LogWriter = (*multiplexedLogWriter)(nil)

// Debug logs a debug message.
func (s multiplexedLogWriter) Debug(msg string) {
	for _, logWriter := range s.logWriters {
		logWriter.Debug(msg)
	}
}

// Info logs an informational message.
func (s multiplexedLogWriter) Info(msg string) {
	for _, logWriter := range s.logWriters {
		logWriter.Info(msg)
	}
}

// Warn logs a warning message.
func (s multiplexedLogWriter) Warn(msg string) {
	for _, logWriter := range s.logWriters {
		logWriter.Warn(msg)
	}
}

// Error logs an error message.
func (s multiplexedLogWriter) Error(msg string) {
	for _, logWriter := range s.logWriters {
		logWriter.Error(msg)
	}
}

// NewLogger creates a Logger from a LogWriter.
func NewLogger(writer LogWriter) Logger {
	return loggerf{writer}
}

// loggerf logs formatted messages to a LogWriter.
type loggerf struct {
	LogWriter
}

var _ Logger = (*loggerf)(nil)

// Debugf logs a debug message.
func (l loggerf) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

// Infof logs an informational message.
func (l loggerf) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

// Warnf logs a warning message.
func (l loggerf) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

// Errorf logs an error message.
func (l loggerf) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

// noopLogger is a Logger that logs nothing.
type noopLogger struct{}

var _ Logger = (*noopLogger)(nil)

// Debug logs nothing.
func (n noopLogger) Debug(msg string) {}

// Debugf logs nothing.
func (n noopLogger) Debugf(format string, args ...interface{}) {}

// Info logs nothing.
func (n noopLogger) Info(msg string) {}

// Infof logs nothing.
func (n noopLogger) Infof(format string, args ...interface{}) {}

// Warn logs nothing.
func (n noopLogger) Warn(msg string) {}

// Warnf logs nothing.
func (n noopLogger) Warnf(format string, args ...interface{}) {}

// Error logs nothing.
func (n noopLogger) Error(msg string) {}

// Errorf logs nothing.
func (n noopLogger) Errorf(format string, args ...interface{}) {}
