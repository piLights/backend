package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/piLights/dioder"
	"github.com/vaughan0/go-ini"
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

	if *configurationFile != "" {
		parseConfiguration(*configurationFile)
	}

	//Set the pins
	if *debug {
		log.Printf("Configuring the Pins to: Red: %d, Green: %d, Blue: %d\n", *redPin, *greenPin, *bluePin)
	}
	dioder.SetPins(*redPin, *greenPin, *bluePin)

	startServer()
}

func parseConfiguration(configurationFile string) {
	//Open the file
	//Read the values for RGB
	//Bindto
	if *debug {
		log.Printf("Parsing configurationFile from %s\n", configurationFile)
	}
	file, error := ini.LoadFile(configurationFile)
	if error != nil {
		log.Fatal(error)
	}

	redPinString, ok := file.Get("PinConfiguration", "RedPin")
	if !ok {
		if *debug {
			log.Println("Value RedPin not set, using default one")
		}
	} else {
		*redPin, _ = strconv.Atoi(redPinString)
	}

	greenPinString, ok := file.Get("PinConfiguration", "GreenPin")
	if !ok {
		if *debug {
			log.Println("Value GreenPin not set, using default one")
		}
	} else {
		*greenPin, _ = strconv.Atoi(greenPinString)
	}

	bluePinString, ok := file.Get("PinConfiguration", "BluePin")
	if !ok && *debug {
		if *debug {
			log.Println("Value BluePin not set, using default one")
		}
	} else {
		*bluePin, _ = strconv.Atoi(bluePinString)
	}

}
