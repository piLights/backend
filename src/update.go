package main

import (
	"fmt"
	"net/http"

	"github.com/inconshreveable/go-update"
)

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
