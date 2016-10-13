package webpage

import (
	"fmt"
	"image/color"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func rgbColorController(w http.ResponseWriter, r *http.Request) {
	redValue := mux.Vars(r)["red"]
	greenValue := mux.Vars(r)["green"]
	blueValue := mux.Vars(r)["blue"]

	responseColorString := redValue + ":" + greenValue + ":" + blueValue
	responseMessage := colorResponse{"Color successfully activated", responseColorString}

	message, error := GenerateResponseMessage(responseMessage)

	if error != nil {
		fmt.Fprintf(w, "There was an error")

		return
	}

	redUint, _ := strconv.Atoi(redValue)
	greenUint, _ := strconv.Atoi(greenValue)
	blueUint, _ := strconv.Atoi(blueValue)

	colorSet := color.RGBA{
		R: uint8(redUint),
		G: uint8(greenUint),
		B: uint8(blueUint),
	}

	dioderConfiguration.DioderInstance.SetAll(colorSet)
	fmt.Fprintf(w, message)
}
