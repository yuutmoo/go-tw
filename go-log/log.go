package log

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type Loglevel string

const (
	DEBUG   Loglevel = "DEBUG"
	INFO    Loglevel = "INFO"
	WARNING Loglevel = "WARNING"
	ERROR   Loglevel = "ERROR"
	PANIC   Loglevel = "PANIC"
	SUCCESS Loglevel = "SUCCESS"
)

var (
	levelFontColors = map[string]string{
		"DEBUG":   "\033[36m", // 青色
		"INFO":    "\033[37m", // 白色
		"WARNING": "\033[33m", // 黄色
		"ERROR":   "\033[31m", // 红色
		"PANIC":   "\033[35m", // 紫色
		"SUCCESS": "\033[32m", // 绿色
	}
	currentLevel = DEBUG

	levelPriority = map[Loglevel]int{
		DEBUG:   1,
		INFO:    2,
		WARNING: 3,
		ERROR:   4,
		PANIC:   5,
		SUCCESS: 6,
	}

	levelColors = map[string]string{
		"DEBUG":   "\033[97;46m",
		"INFO":    "\033[30;47m",
		"WARNING": "\033[30;43m",
		"ERROR":   "\033[97;41m",
		"PANIC":   "\033[97;45m",
		"SUCCESS": "\033[1;97;42m",
	}

	logOutputs = []logOutput{
		{writer: os.Stdout, isConsole: true},
	}
	lastLogFileDate = ""
	logFile         *os.File
)

type logOutput struct {
	writer    *os.File
	isConsole bool
}

func AddLogOutputFile(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	logOutputs = append(logOutputs, logOutput{writer: file, isConsole: false})
	return nil
}

func ensureLogFile() {
	today := time.Now().Format("2006-01-02")
	if lastLogFileDate == today && logFile != nil {
		return
	}
	if logFile != nil {
		logFile.Close()
		logFile = nil
	}
	_ = os.MkdirAll("log", 0755)
	logPath := filepath.Join("log", today+".log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err == nil {
		logFile = file
		found := false
		for _, out := range logOutputs {
			if out.writer == file {
				found = true
				break
			}
		}
		if !found {
			logOutputs = append(logOutputs, logOutput{writer: file, isConsole: false})
		}
		lastLogFileDate = today
	}
}

func writeLogAllTargets(coloredMsg string, plainMsg string) {
	ensureLogFile()
	for _, out := range logOutputs {
		if out.isConsole {
			fmt.Fprintln(out.writer, coloredMsg)
		} else {
			fmt.Fprintln(out.writer, plainMsg)
		}
	}
}

func SetLogLevel(level Loglevel) {
	currentLevel = level
}

func shouldLog(level Loglevel) bool {
	return levelPriority[level] >= levelPriority[currentLevel]
}

func formatLog(level string, message string, withColor bool) string {
	now := time.Now().Format("2006-01-02 15:04:05")
	pc, file, line, ok := runtime.Caller(3)
	caller := "unknown"
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		caller = fmt.Sprintf("%s:%d", file, line)
		if index := strings.LastIndex(funcName, "."); index != -1 {
			caller = fmt.Sprintf("%s.%s:%d", funcName[:index], funcName[index+1:], line)
		}
	}

	levelStr := fmt.Sprintf("%-8s", level)

	if withColor {
		return fmt.Sprintf("\033[32m%s\033[0m |\033[0m%s%s\033[0m| \033[35m%s\033[0m - %s%s\033[0m",
			now, levelColors[level], levelStr, caller, levelFontColors[level], message) // 使用 levelFontColors 设置消息颜色
	}
	return fmt.Sprintf("%s |%s| %s - %s",
		now, levelStr, caller, message)
}

func logWithLevel(level string, v ...any) {
	lv := Loglevel(level)
	if !shouldLog(lv) {
		return
	}

	var parts []string
	for _, arg := range v {
		val := reflect.ValueOf(arg)
		kind := val.Kind()
		if kind == reflect.Ptr {
			if val.IsNil() {
				parts = append(parts, "<nil>")
				continue
			}
			val = val.Elem()
			kind = val.Kind()
		}

		if err, ok := arg.(error); ok {
			parts = append(parts, err.Error())
		} else if kind == reflect.Struct || kind == reflect.Map || kind == reflect.Array {
			jsonData, err := json.MarshalIndent(arg, "", "  ")
			if err == nil {
				parts = append(parts, val.Type().Name(), " "+string(jsonData))
			} else {
				parts = append(parts, fmt.Sprint(arg))
			}
		} else {
			parts = append(parts, fmt.Sprint(arg))
		}
	}
	message := strings.Join(parts, " ")

	coloredLogLine := formatLog(level, message, true)
	plainLogLine := formatLog(level, message, false)
	writeLogAllTargets(coloredLogLine, plainLogLine)
}

func logWithLevelf(level string, format string, v ...any) {
	lv := Loglevel(level)
	if !shouldLog(lv) {
		return
	}
	message := fmt.Sprintf(format, v...)
	coloredLogLine := formatLog(level, message, true)
	plainLogLine := formatLog(level, message, false)
	writeLogAllTargets(coloredLogLine, plainLogLine)
}


type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Print(message string) {
	logWithLevel(string(INFO), message)
}

func (l *Logger) Trace(message string) {
	logWithLevel(string(DEBUG), message)
}

func (l *Logger) Debug(message string) {
	logWithLevel(string(DEBUG), message)
}

func (l *Logger) Info(message string) {
	logWithLevel(string(INFO), message)
}

func (l *Logger) Warning(message string) {
	logWithLevel(string(WARNING), message)
}

func (l *Logger) Error(message string) {
	logWithLevel(string(ERROR), message)
}

func (l *Logger) Fatal(message string) {
	logWithLevel(string(PANIC), message)
	os.Exit(1)
}

func Print(v ...any) {
	logWithLevel(string(INFO), v...)
}

func Printf(format string, v ...any) {
	logWithLevelf("INFO", format, v...)
}

func Println(v ...any) {
	logWithLevel(string(INFO), v...)
}

func Fatal(v ...any) {
	logWithLevel(string(ERROR), v...)
	os.Exit(1)
}

func Fatalf(format string, v ...any) {
	logWithLevelf(string(ERROR), format, v...)
	os.Exit(1)
}

func Fatalln(v ...any) {
	logWithLevel(string(ERROR), v...)
	os.Exit(1)
}

func Panic(v ...any) {
	s := fmt.Sprint(v...)
	logWithLevel(string(PANIC), s)
	panic(s)
}

func Panicf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	logWithLevel(string(PANIC), s)
	panic(s)
}

func Panicln(v ...any) {
	s := fmt.Sprintln(v...)
	logWithLevel(string(PANIC), s)
	panic(s)
}

func Debug(v ...any)    { logWithLevel("DEBUG", v...) }
func Info(v ...any)     { logWithLevel(string(INFO), v...) }
func Warning(v ...any)  { logWithLevel("WARNING", v...) }
func Error(v ...any)    { logWithLevel(string(ERROR), v...) }
func Critical(v ...any) { logWithLevel("PANIC", v...) }
func Success(v ...any)  { logWithLevel("SUCCESS", v...) }

func Debugf(format string, v ...any)    { logWithLevelf(string(DEBUG), format, v...) }
func Infof(format string, v ...any)     { logWithLevelf("INFO", format, v...) }
func Warningf(format string, v ...any)  { logWithLevelf(string(WARNING), format, v...) }
func Errorf(format string, v ...any)    { logWithLevelf(string(ERROR), format, v...) }
func Criticalf(format string, v ...any) { logWithLevelf(string(PANIC), format, v...) }
func Successf(format string, v ...any)  { logWithLevelf(string(SUCCESS), format, v...) }
