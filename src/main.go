package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/piLights/dioder"
	"github.com/piLights/dioder-rpc/src/configuration"
	"github.com/piLights/dioder-rpc/src/logging"
	"github.com/piLights/dioder-rpc/src/server"
	"github.com/urfave/cli"
)

const version = "debugVersion"

func main() {
	application := cli.NewApp()
	application.Name = "Dioder-Server"
	application.Version = version
	application.Flags = applicationFlags

	application.Action = func(c *cli.Context) error {
		logging.NewLoggingService()
		go logging.Log.Start()

		//Check, if we should update
		if c.Bool("update") {
			startUpdate()
			return nil
		}

		if c.String("writeConfiguration") != "" {
			error := configuration.DioderConfiguration.WriteConfigurationToFile(c.String("writeConfiguration"))

			return error
		}

		if c.String("configurationFile") != "" {
			error := configuration.NewConfiguration(c.String("configurationFile"))

			if error != nil {
				return error
			}
		}

		//Set the pins
		if configuration.DioderConfiguration.Debug {
			logging.Log.LogChan <- fmt.Sprintf("Configuring the Pins to: Red: %s, Green: %s, Blue: %s", configuration.DioderConfiguration.Pins.Red, configuration.DioderConfiguration.Pins.Green, configuration.DioderConfiguration.Pins.Blue)
		}

		go startAutoConfigurationServer()

		dioderInstance := dioder.New(dioder.Pins{Red: configuration.DioderConfiguration.Pins.Red, Green: configuration.DioderConfiguration.Pins.Green, Blue: configuration.DioderConfiguration.Pins.Blue}, configuration.DioderConfiguration.PiBlaster)

		server.StartServer(dioderInstance)

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
