package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/piLights/dioder"
	"github.com/urfave/cli"
)

const version = "debugVersion"

var (
	dioderInstance dioder.Dioder
	logChan        chan interface{}
	fatalChan      chan interface{}
)

func main() {
	application := cli.NewApp()
	application.Name = "Dioder-Server"
	application.Version = version
	application.Flags = applicationFlags

	application.Action = func(c *cli.Context) error {
		logChan = make(chan interface{}, 100)
		fatalChan = make(chan interface{}, 100)
		go loggingService(logChan, fatalChan)

		//Check, if we should update
		if doUpdate {
			startUpdate()
			return nil
		}

		/*if *configurationFile != "" {
			parseConfiguration(DioderConfiguration.ConfigurationFile)
		}*/

		//Set the pins
		if DioderConfiguration.Debug {
			logChan <- fmt.Sprintf("Configuring the Pins to: Red: %s, Green: %s, Blue: %s", DioderConfiguration.Pins.Red, DioderConfiguration.Pins.Green, DioderConfiguration.Pins.Blue)
		}

		go startAutoConfigurationServer()

		dioderInstance = dioder.New(dioder.Pins{Red: DioderConfiguration.Pins.Red, Green: DioderConfiguration.Pins.Green, Blue: DioderConfiguration.Pins.Blue}, DioderConfiguration.PiBlaster)

		startServer()

		return nil
	}

	application.Run(os.Args)

}

//hashPassword hashes the defined password with SHA256
func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))

	return hex.EncodeToString(hasher.Sum(nil))
}
