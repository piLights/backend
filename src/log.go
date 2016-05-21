package main

import (
	"log"
	"os"
)

func loggingService(logChan, fatalChan chan interface{}) {
	//Ldate | Ltime | Lmicroseconds | Llongfile
	//Open

	if *debug {
		logChan <- "Starting the logging service."
	}

	logInstance := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LstdFlags)

	for {
		select {
		case logLine := <-logChan:
			logInstance.Println(logLine)
		case failureLogLine := <-fatalChan:
			logInstance.Fatalln(failureLogLine)
		}
	}
}
