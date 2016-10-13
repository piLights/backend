package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"gitlab.com/piLights/proto"
)

const maxSliceSize = math.MaxInt32

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
	// Check if the push of another entry would cause the EntryList-Slice to overflow
	if logList.Count+1 > maxSliceSize {
		// Remove the first item
		logList.EntryList = append(logList.EntryList[:0], logList.EntryList[1:]...)
		logList.Count--
	}

	entry := &LighterGRPC.LogEntry{
		Time:    time.Now().UnixNano(),
		Message: fmt.Sprintf("%v", line), // Fix: Type assertions failing here: We have string and *grpc.error
	}

	logList.EntryList = append(logList.EntryList, entry)
	logList.Count++
}
