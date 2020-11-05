package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup() {
	var err error
	fileName := getLogFileName()
	filePath := getLogFilePath()
	F, err = openLogFile(fileName, filePath)
	if err != nil {
		log.Fatalln(err)
	}
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

//控制台和文件均会输出
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
	log.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
	log.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
	log.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
	log.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
	log.SetPrefix(logPrefix)
}

func LineInfo() string {
	function := "xxx"
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	function = runtime.FuncForPC(pc).Name()
	return fmt.Sprintf(" -> %s():%s:%d", function, file, line)
}
