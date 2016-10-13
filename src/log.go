package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gitlab.com/piLights/proto"
)

type logEntryList struct {
	EntryList []*LighterGRPC.LogEntry
	Count     int
}

var logList logEntryList

func getLogEntryList() logEntryList {
	return logList
}

func loggingService(logChan, fatalChan chan interface{}) {
	//Ldate | Ltime | Lmicroseconds | Llongfile
	//Open

	if DioderConfiguration.Debug {
		logChan <- "Starting the logging service."
	}

	logInstance := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LstdFlags)

	for {

		select {
		case logLine := <-logChan:
			saveLog(logLine)
			logInstance.Println(logLine)
		case failureLogLine := <-fatalChan:
			saveLog(failureLogLine)
			logInstance.Fatalln(failureLogLine)
		}
	}
}

func saveLog(line interface{}) {
	entry := &LighterGRPC.LogEntry{
		Time:    time.Now().UnixNano(),
		Message: fmt.Sprintf("%v", line), // Fix: Type assertions failing here: We have string and *grpc.error
	}

	logList.EntryList = append(logList.EntryList, entry)
	logList.Count++
}
