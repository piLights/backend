package logging

import (
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"time"

	"gitlab.com/piLights/dioder-rpc/configuration"
	"gitlab.com/piLights/proto"
)

const maxSliceSize = math.MaxInt32

type logEntryList struct {
	EntryList []*LighterGRPC.LogEntry
	Count     int32
	mutex     sync.Mutex
}

var logList logEntryList

var (
	// FatalChan is the channel used for fatal errors
	FatalChan chan interface{}

	// LogChan is the channel used for anything related to logging
	LogChan chan interface{}
)

// GetLogEntryList fetches the list of LogEntries
func GetLogEntryList(amount int32) []*LighterGRPC.LogEntry {
	logList.mutex.Lock()
	defer logList.mutex.Unlock()
	if amount >= logList.Count {
		return logList.EntryList
	}

	return logList.EntryList[logList.Count-amount:]
}

// NewLoggingService instanciates a new logging service
func NewLoggingService() {
	FatalChan = make(chan interface{}, 100)
	LogChan = make(chan interface{}, 100)
}

func Service() {
	//Ldate | Ltime | Lmicroseconds | Llongfile
	//Open

	if configuration.DioderConfiguration.Debug {
		LogChan <- "Starting the logging service."
	}

	logInstance := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LstdFlags)

	for {

		select {
		case logLine := <-LogChan:
			saveLog(logLine)
			logInstance.Println(logLine)
		case failureLogLine := <-FatalChan:
			saveLog(failureLogLine)
			logInstance.Fatalln(failureLogLine)
		}
	}
}

func saveLog(line interface{}) {
	logList.mutex.Lock()
	defer logList.mutex.Unlock()
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
