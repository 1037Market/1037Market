package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Logger interface {
	Info(message string)
	Error(message string)
	Fatal(message string)
	Close()
}

type FileLogger struct {
	file *os.File
}

func (l *FileLogger) Info(v ...interface{}) {
	message := fmt.Sprintln(v...)
	fmt.Fprintf(l.file, "INFO: %s\n", message)
}

func (l *FileLogger) Error(v ...interface{}) {
	message := fmt.Sprintln(v...)
	fmt.Fprintf(l.file, "Error: %s\n", message)
}

func (l *FileLogger) Fatal(v ...interface{}) {
	message := fmt.Sprintln(v...)
	fmt.Fprintf(l.file, "Fatal: %s\n", message)
}

var (
	instance *FileLogger
	once     sync.Once
)

func GetInstance() *FileLogger {
	once.Do(func() {
		file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("Failed to open log file:", err)
		}
		instance = &FileLogger{
			file: file,
		}
	})
	return instance
}
