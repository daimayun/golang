package logs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Level int

var (
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	levelFlags         = map[Level]string{DEBUG: "DEBUG", INFO: "INFO", WARNING: "WARN", ERROR: "ERROR", FATAL: "FATAL", ORM: "ORM"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
	ORM
)

func init() {
	for _, label := range levelFlags {
		createLogFile(getLogFileFullPath(strings.ToLower(label)))
	}
}

func NewLogger(f *os.File) *log.Logger {
	return log.New(f, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	writeLog(DEBUG, v...)
}

func Info(v ...interface{}) {
	writeLog(INFO, v...)
}

func Warn(v ...interface{}) {
	writeLog(WARNING, v...)
}

func Warning(v ...interface{}) {
	Warn(v...)
}

func Error(v ...interface{}) {
	writeLog(ERROR, v...)
}

func Fatal(v ...interface{}) {
	f := openLogFile(getLogFileFullPath(strings.ToLower(levelFlags[FATAL])))
	defer f.Close()
	logger := NewLogger(f)
	logger.SetPrefix(getPrefix(FATAL))
	logger.Fatalln(v)
}

func Orm(v ...interface{}) {
	writeLog(ORM, v...)
}

func writeLog(level Level, v ...interface{}) {
	f := openLogFile(getLogFileFullPath(strings.ToLower(levelFlags[level])))
	defer f.Close()
	logger := NewLogger(f)
	logger.SetPrefix(getPrefix(level))
	logger.Println(v)
}

func getPrefix(level Level) string {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		return fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	}
	return fmt.Sprintf("[%s]", levelFlags[level])
}
