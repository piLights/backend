package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/inconshreveable/go-update"
)

const UPDATEURL = "https://github.com/piLights/dioder-rpc/releases/download/pre-release/dioderAPI_" + runtime.GOOS + "_" + runtime.GOARCH + "_src"

func startUpdate() {
	fmt.Println("Starting update...")
	if *debug {
		log.Printf("Downloading from %s\n", *updateFromURL)
	}
	error := updateBinary(*updateFromURL)
	if error != nil {
		fmt.Println("Updating failed!")
		log.Fatal(error)
	}

	fmt.Println("Updated successfully")
}

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
