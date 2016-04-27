package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"runtime/pprof"
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
		log.Printf("Configuring the Pins to: Red: %d, Green: %d, Blue: %d\n", *redPin, *greenPin, *bluePin)
	}
	dioder.SetPins(*redPin, *greenPin, *bluePin)

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

	passwordString, ok := file.Get("General", "Password")
	if !ok && *debug {
		if *debug {
			log.Println("Value Password not set, using none")
		}
	} else {
		*password = passwordString
	}

}
