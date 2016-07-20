package logging

import (
	"log"
	"os"

	"github.com/piLights/dioder-rpc/src/configuration"
)

type Logger struct {
	LogChan     chan interface{}
	FatalChan   chan interface{}
	logInstance *log.Logger
}

var Log Logger

func NewLoggingService() Logger {
	logChan := make(chan interface{}, 100)
	fatalChan := make(chan interface{}, 100)
	logInstance := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LstdFlags)

	Log := Logger{
		LogChan:     logChan,
		FatalChan:   fatalChan,
		logInstance: logInstance,
	}

	return Log
}

func (logger *Logger) Start() {
	if configuration.DioderConfiguration.Debug {
		logger.LogChan <- "Starting the logging service."
	}

	for {
		select {
		case logLine := <-logger.LogChan:
			logger.logInstance.Println(logLine)
		case failureLogLine := <-logger.FatalChan:
			logger.logInstance.Fatalln(failureLogLine)
		}
	}

}
