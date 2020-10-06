// An extension to go-kit framework logger to add log level and initializer with configuration
package logger

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
)

var (
	loggerInstance log.Logger
	once           sync.Once
	syncWriter     io.Writer
)

type CustomLogger struct {
	logger log.Logger
}


type Conf struct {
	LogLevel   string // levels: NONE|ERROR|WARN|INFO|DEBUG|ALL
	TZLocal    bool   // true: local timezone else UTC
}

// Initializes a logger with the given configuration
func InitLogger(conf Conf) {
	once.Do(func() {
		var logWriter io.Writer
		logWriter = os.Stdout
		syncWriter = log.NewSyncWriter(logWriter)
		loggerInstance = log.NewLogfmtLogger(syncWriter)
		if conf.TZLocal {
			loggerInstance = log.With(loggerInstance, "time", log.DefaultTimestamp)
		} else {
			loggerInstance = log.With(loggerInstance, "time", log.DefaultTimestampUTC)
		}
		loggerInstance = level.NewFilter(loggerInstance, getLogLevel(conf.LogLevel))

	})
}

// Returns Logger anytime after initialization
func GetLogger() CustomLogger {
	return CustomLogger{
		logger: loggerInstance,
	}
}

func getCallerInfo() string {
	var callerInfo string
	if _, file, line, ok := runtime.Caller(2); ok {
		file = path.Base(file)
		callerInfo = fmt.Sprintf("%s:%d", file, line)
	}
	return callerInfo
}

// Returns level for the given string level
func getLogLevel(levelStr string) level.Option {
	switch strings.ToUpper(levelStr) {
	case "INFO":
		return level.AllowInfo()
	case "DEBUG":
		return level.AllowDebug()
	case "WARN":
		return level.AllowWarn()
	case "ERROR":
		return level.AllowError()
	case "NONE":
		return level.AllowNone()
	case "ALL":
		return level.AllowAll()
	default:
		return level.AllowInfo()
	}
}

// Logs with No level.
func (cLogger CustomLogger) Log(args ...interface{}) error {
	return log.With(loggerInstance, "caller", getCallerInfo()).Log(args...)
}

// Logs with Info level.
func (cLogger CustomLogger) Info(args ...interface{}) {
	level.Info(log.With(loggerInstance, "caller", getCallerInfo())).Log(args...)
}

// Logs with Debug level.
func (cLogger CustomLogger) Debug(args ...interface{}) {
	level.Debug(log.With(loggerInstance, "caller", getCallerInfo())).Log(args...)
}

// Logs with Error level.
func (cLogger CustomLogger) Error(args ...interface{}) {
	level.Error(log.With(loggerInstance, "caller", getCallerInfo())).Log(args...)
	printStackTraceIfPresent(args...)
}

// Logs with Warn level.
func (cLogger CustomLogger) Warn(args ...interface{}) {
	level.Warn(log.With(loggerInstance, "caller", getCallerInfo())).Log(args...)
}

// Logs with Info level.
func (cLogger CustomLogger) Infom(arg interface{}) {
	level.Info(log.With(loggerInstance, "caller", getCallerInfo())).Log("message", arg)
}

func (cLogger CustomLogger) Debugm(arg interface{}) {
	level.Debug(log.With(loggerInstance, "caller", getCallerInfo())).Log("message", arg)
}

// Logs with Error level.
func (cLogger CustomLogger) Errorm(arg interface{}) {
	level.Error(log.With(loggerInstance, "caller", getCallerInfo())).Log("error", arg)
	printStackTraceIfPresent(arg)
}


// Logs with Info level.
func (cLogger CustomLogger) Infof(format string, args ...interface{}) {
	level.Info(log.With(loggerInstance, "caller", getCallerInfo())).Log("message", fmt.Sprintf(format, args...))
}

func (cLogger CustomLogger) Errorf(format string, args ...interface{}) {
	level.Error(log.With(loggerInstance, "caller", getCallerInfo())).Log("error", fmt.Errorf(format, args...))
	printStackTraceIfPresent(args...)
}

// prints stacktrace to the log.
func printStackTraceIfPresent(args ...interface{}) {
	for i := 0; i < len(args); i += 1 {
		switch value := args[i].(type) {
		// If an error type and has a stacktrace then log it.
		case error:
			if syncWriter != nil {
				if err, ok := value.(stackTracer); ok {
					fmt.Fprintf(syncWriter, "%+v\n", err)
				}
			}
		default:
			// Ignore all other types
		}
	}
}

// Interface to retrieve stack trace of an error
type stackTracer interface {
	StackTrace() errors.StackTrace
}
