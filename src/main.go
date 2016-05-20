package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/piLights/dioder"
	"github.com/vaughan0/go-ini"
)

const version = "debugVersion"

var (
	dioderInstance dioder.Dioder
	logChan        chan interface{}
	fatalChan      chan interface{}
)

func main() {
	dioderAPI.Version(version)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.MustParse(dioderAPI.Parse(os.Args[1:]))

	logChan = make(chan interface{}, 100)
	fatalChan = make(chan interface{}, 100)
	go loggingService(logChan, fatalChan)

	//Only for debugging!
	if *cpuProfile != "" {
		file, error := os.Create(*cpuProfile)
		if error != nil {
			log.Fatal(error)
		}

		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}

	//Check, if we should update
	if *doUpdate {
		startUpdate()
		return
	}

	if *configurationFile != "" {
		parseConfiguration(*configurationFile)
	}

	if *password != "" {
		*password = hashPassword(*password)
	}

	//Set the pins
	if *debug {
		logChan <- fmt.Sprintf("Configuring the Pins to: Red: %s, Green: %s, Blue: %s", *redPin, *greenPin, *bluePin)
	}

	go startAutoConfigurationServer()

	dioderInstance = dioder.New(dioder.Pins{*redPin, *greenPin, *bluePin}, *piBlaster)

	startServer()
}

//hashPassword hashes the defined password with SHA256
func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))

	return hex.EncodeToString(hasher.Sum(nil))
}

//parseConfiguration parses the configurationFile and sets the specified values
func parseConfiguration(configurationFile string) {
	//Open the file
	//Read the values for RGB
	//Bindto
	if *debug {
		logChan <- fmt.Sprintf("Parsing configurationFile from %s", configurationFile)
	}
	file, error := ini.LoadFile(configurationFile)
	if error != nil {
		log.Fatal(error)
	}

	redPinString, ok := file.Get("PinConfiguration", "RedPin")
	if !ok {
		logChan <- fmt.Sprintf("Value RedPin not set, using default: %s", *redPin)
	} else {
		*redPin = redPinString
	}

	greenPinString, ok := file.Get("PinConfiguration", "GreenPin")
	if !ok {
		logChan <- fmt.Sprintf("Value GreenPin not set, using default %s", *greenPin)
	} else {
		*greenPin = greenPinString
	}

	bluePinString, ok := file.Get("PinConfiguration", "BluePin")
	if !ok && *debug {
		logChan <- fmt.Sprintf("Value BluePin not set, using default %s", *bluePin)
	} else {
		*bluePin = bluePinString
	}

	passwordString, ok := file.Get("General", "Password")
	if !ok && *debug {
		logChan <- fmt.Sprintf("Value Password not set, using none")
	} else {
		*password = passwordString
	}

	piBlasterLocation, ok := file.Get("General", "PiBlaster")
	if !ok && *debug {
		logChan <- fmt.Sprintf("Value PiBlaster not set, using default: %s", *piBlaster)
	} else {
		*piBlaster = piBlasterLocation
	}

	if *piBlaster == "" {
		*piBlaster = "/dev/pi-blaster"
	}

	configServerName, ok := file.Get("General", "ServerName")
	if !ok && *debug {
		logChan <- fmt.Sprintf("Value ServerName not set, using default: %s", *piBlaster)
	} else {
		*serverName = configServerName
	}

	if *serverName == "" {
		*serverName = "Dioder Server"
	}
}
