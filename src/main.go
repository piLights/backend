package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/signal"

	"gitlab.com/piLights/dioder"
	"gitlab.com/piLights/dioder-rpc/src/configuration"
	"gitlab.com/piLights/dioder-rpc/src/logging"
	"gitlab.com/piLights/dioder-rpc/src/piLightsVersion"
	"gitlab.com/piLights/dioder-rpc/src/rpc"

	"github.com/urfave/cli"
)

func main() {
	// Report crashes
	//defer captureCrash()

	application := cli.NewApp()
	application.Name = "Dioder-Server"
	application.Version = piLightsVersion.Version
	application.Flags = applicationFlags

	application.Action = func(c *cli.Context) error {
		logging.NewLoggingService()
		go logging.Service()

		//Check, if we should update
		if c.Bool("update") {
			startUpdate()
			return nil
		}

		if c.String("writeConfiguration") != "" {
			fmt.Println(configuration.DioderConfiguration)
			error := configuration.DioderConfiguration.WriteConfigurationToFile(c.String("writeConfiguration"))

			return error
		}

		if c.String("configurationFile") != "" {
			config, err := configuration.NewConfiguration(c.String("configurationFile"))
			if err != nil {
				return err
			}

			configuration.DioderConfiguration = config
		}

		//Set the pins
		if configuration.DioderConfiguration.Debug {
			logging.LogChan <- fmt.Sprintf("Configuring the Pins to: Red: %s, Green: %s, Blue: %s", configuration.DioderConfiguration.Pins.Red, configuration.DioderConfiguration.Pins.Green, configuration.DioderConfiguration.Pins.Blue)
		}

		if !configuration.DioderConfiguration.NoAutoconfiguration {
			go startAutoConfigurationServer()
		}

		configuration.DioderConfiguration.DioderInstance = dioder.New(dioder.Pins{Red: configuration.DioderConfiguration.Pins.Red, Green: configuration.DioderConfiguration.Pins.Green, Blue: configuration.DioderConfiguration.Pins.Blue}, configuration.DioderConfiguration.PiBlaster)

		//Handle CTRL  + C
		osSignalChan := make(chan os.Signal, 1)
		signal.Notify(osSignalChan, os.Interrupt)
		go func() {
			<-osSignalChan

			// Close channels
			close(osSignalChan)
			close(logging.LogChan)
			close(logging.FatalChan)

			// Release all Pins for further use
			configuration.DioderConfiguration.DioderInstance.Release()

			os.Exit(0)
		}()

		rpc.StartServer()

		defer configuration.DioderConfiguration.DioderInstance.Release()

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
