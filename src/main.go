package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/signal"

	"gitlab.com/piLights/dioder-rpc/src/configuration"
	"gitlab.com/piLights/dioder-rpc/src/webpage"

	"github.com/piLights/dioder"
	"github.com/urfave/cli"
)

const version = "debugVersion"

var (
	logChan   chan interface{}
	fatalChan chan interface{}
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
		if c.Bool("update") {
			startUpdate()
			return nil
		}

		if c.String("writeConfiguration") != "" {
			fmt.Println(DioderConfiguration)
			error := DioderConfiguration.WriteConfigurationToFile(c.String("writeConfiguration"))

			return error
		}

		if c.String("configurationFile") != "" {
			config, err := configuration.NewConfiguration(c.String("configurationFile"))
			if err != nil {
				return err
			}

			DioderConfiguration = config
		}

		//Set the pins
		if DioderConfiguration.Debug {
			logChan <- fmt.Sprintf("Configuring the Pins to: Red: %s, Green: %s, Blue: %s", DioderConfiguration.Pins.Red, DioderConfiguration.Pins.Green, DioderConfiguration.Pins.Blue)
		}

		if !DioderConfiguration.NoAutoconfiguration {
			go startAutoConfigurationServer()
		}

		DioderConfiguration.DioderInstance = dioder.New(dioder.Pins{Red: DioderConfiguration.Pins.Red, Green: DioderConfiguration.Pins.Green, Blue: DioderConfiguration.Pins.Blue}, DioderConfiguration.PiBlaster)

		//Handle CTRL  + C
		osSignalChan := make(chan os.Signal, 1)
		signal.Notify(osSignalChan, os.Interrupt)
		go func() {
			<-osSignalChan

			// Close channels
			close(colorStream)
			close(osSignalChan)
			close(logChan)
			close(fatalChan)

			// Release all Pins for further use
			DioderConfiguration.DioderInstance.Release()

			os.Exit(0)
		}()

		if DioderConfiguration.Debug {
			go webpage.StartWebPage(logChan, fatalChan, DioderConfiguration)
		}

		startServer()

		defer DioderConfiguration.DioderInstance.Release()

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
