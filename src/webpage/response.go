package webpage

import (
	"encoding/json"
)

type colorResponse struct {
	Message string
	Color   string
}

func GenerateResponseMessage(v interface{}) (string, error) {
	message, error := json.Marshal(v)

	if error != nil {
		return "", error
	}

	return string(message), nil
}
