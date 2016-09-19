package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/oleksandr/bonjour"
)

//startServer starts the GRPC-server and binds to the defined address
func startAutoConfigurationServer() {
	if DioderConfiguration.Debug {
		logChan <- fmt.Sprintf("Binding to %s", DioderConfiguration.BindTo)
	}

	/*_, port, error := net.SplitHostPort(DioderConfiguration.BindTo)
	if error != nil {
		log.Fatal(error)
	}*/

	s, err := bonjour.Register("dioderServer", "_dioder._tcp", "", 13337, nil, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Ctrl+C handling
	handler := make(chan os.Signal, 1)
	signal.Notify(handler, os.Interrupt)
	for sig := range handler {
		if sig == os.Interrupt {
			s.Shutdown()
			time.Sleep(1e9)
			break
		}
	}
}
