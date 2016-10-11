package main

import (
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
		Message: line.(string), // @ToDo: Maybe fmt.Sprintf(line) would be better
	}

	logList.EntryList = append(logList.EntryList, entry)
	logList.Count++
}
