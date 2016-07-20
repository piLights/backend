package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/inconshreveable/go-update"
	"github.com/piLights/dioder-rpc/src/configuration"
	"github.com/piLights/dioder-rpc/src/logging"
)

//UPDATEURL is the URL from which the updates for the OS and architecture are fetched from
const UPDATEURL = "https://github.com/piLights/dioder-rpc/releases/download/pre-release/dioderAPI_" + runtime.GOOS + "_" + runtime.GOARCH

//startUpdate starts the updateProcess
func startUpdate() {
	fmt.Println("Starting update...")
	if configuration.DioderConfiguration.Debug {
		logging.Log.LogChan <- fmt.Sprintf("Downloading from %s\n", configuration.DioderConfiguration.UpdateURL)
	}
	error := updateBinary(configuration.DioderConfiguration.UpdateURL)
	if error != nil {
		fmt.Println("Updating failed!")
		logging.Log.FatalChan <- error
	}

	fmt.Println("Updated successfully")
}

//updateBinary updates the executable binary
func updateBinary(url string) error {
	// request the new file
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	error := update.Apply(resp.Body, update.Options{})
	if error != nil {
		rollbackError := update.RollbackError(err)
		if rollbackError != nil {
			fmt.Printf("Failed to rollback from bad update: %v\n", rollbackError)
		}
	}

	return error
}
