// Simple log library wrapper - exposes Debug, Info, Error, Warning, Critical
//
// Thanks @ http://www.goinggo.net/2013/11/using-log-package-in-go.html
package golog

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	Debug    *log.Logger
	Info     *log.Logger
	Warning  *log.Logger
	Error    *log.Logger
	Critical *log.Logger
}

func New(filename string, quiet bool) (*Logger, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return &Logger{}, err
	}

	var handle io.Writer

	if quiet {
		handle = file
	} else {
		handle = io.MultiWriter(file, os.Stdout)
	}

	logger := &Logger{}
	logger.SetupLoggers(handle)

	return logger, nil
}

func (l *Logger) SetupLoggers(handle io.Writer) {
	l.Debug = log.New(handle, "DEBUG ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Info = log.New(handle, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Warning = log.New(handle, "WARNING ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Error = log.New(handle, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Critical = log.New(handle, "CRITICAL ", log.Ldate|log.Ltime|log.Lshortfile)
}
