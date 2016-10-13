package webpage

import (
	"fmt"
	"image/color"
	"net/http"

	"github.com/gorilla/mux"
)

func colorController(w http.ResponseWriter, r *http.Request) {
	colorString := mux.Vars(r)["color"]

	responseMessage := colorResponse{"Color successfully activated", colorString}

	message, error := GenerateResponseMessage(responseMessage)

	if error != nil {
		fmt.Fprintf(w, "There was an error")

		return
	}

	fmt.Fprintf(w, message)

	var colorSet color.RGBA

	switch colorString {
	case "red":
		colorSet.R = 255
		break
	case "green":
		colorSet.G = 255
		break
	case "blue":
		colorSet.B = 255
		break
	}

	dioderConfiguration.DioderInstance.SetAll(colorSet)
}
