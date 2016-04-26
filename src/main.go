package main

import (
	"fmt"
	"log"
	"net"
	"runtime"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/piLights/dioder"

	LighterGRPC "github.com/piLights/dioder-rpc/src/proto"
	"google.golang.org/grpc"
)

const version = "debugVersion"

func main() {
	kingpin.Version(version)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	//Check, if we should update
	if *doUpdate {
		//Build url
		updateURL := "https://github.com/piLights/dioder-rpc/releases/download/pre-release/dioderAPI_" + runtime.GOOS + "_" + runtime.GOARCH + "_src"

		if *updateFromURL != "" {
			updateURL = *updateFromURL
		}

		fmt.Println("Starting update...")
		if *debug {
			log.Printf("Downloading from %s\n", updateURL)
		}
		error := updateBinary(updateURL)
		if error != nil {
			fmt.Println("Updating failed!")
			log.Fatal(error)
		}

		fmt.Println("Updated successfully")
		return
	}

	//Set the pins
	if *debug {
		log.Printf("Configuring the Pins to: Red: %d, Green: %d, Blue: %d\n", *redPin, *greenPin, *bluePin)
	}
	dioder.SetPins(*redPin, *greenPin, *bluePin)

	if *debug {
		log.Printf("Binding to %s", *bindTo)
	}
	listener, error := net.Listen("tcp", *bindTo)
	if error != nil {
		log.Fatalf("failed to listen: %v", error)
	}

	grpcServer := grpc.NewServer()

	LighterGRPC.RegisterLighterServer(grpcServer, &server{})

	grpcServer.Serve(listener)
}
