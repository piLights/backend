package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/inconshreveable/go-update"
)

//UPDATEURL is the URL from which the updates for the OS and architecture are fetched from
const UPDATEURL = "http://dl.pilights.jf-projects.de/dioderAPI_" + runtime.GOOS + "_" + runtime.GOARCH

var errFileNoFound = errors.New("File not found")

//startUpdate starts the updateProcess
func startUpdate() {
	fmt.Println("Starting update...")
	if DioderConfiguration.Debug {
		logChan <- fmt.Sprintf("Downloading from %s\n", DioderConfiguration.UpdateURL)
	}
	error := updateBinary(DioderConfiguration.UpdateURL)
	if error != nil {
		fmt.Println("Updating failed!")
		log.Fatal(error)
	}

	fmt.Println("Updated successfully")
}

//updateBinary updates the executable binary
func updateBinary(url string) error {
	// Fetch the Hash-Sum
	/*checksumResponse, error := http.Get(url + ".sha256")
	if error != nil {
		return error
	}
	defer checksumResponse.Body.Close()

	if checksumResponse.StatusCode != 200 {
		return errFileNoFound
	}

	checksum, error := ioutil.ReadAll(checksumResponse.Body)
	if error != nil {
	return error
	}*/

	// request the new file
	response, error := http.Get(url)
	if error != nil {
		return error
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errFileNoFound
	}

	error = update.Apply(response.Body, update.Options{
	//Checksum: checksum,
	})
	if error != nil {
		rollbackError := update.RollbackError(error)
		if rollbackError != nil {
			fmt.Printf("Failed to rollback from bad update: %v\n", rollbackError)
		}
	}

	return error
}
