package webpage

import (
	"fmt"
	"image/color"
	"net/http"

	sadboxColor "gitlab.com/jannickfahlbusch/color"

	"github.com/gorilla/mux"
)

func hexColorController(w http.ResponseWriter, r *http.Request) {
	hexColor := mux.Vars(r)["hex"]

	responseMessage := colorResponse{"Color successfully activated", hexColor}

	message, error := GenerateResponseMessage(responseMessage)

	if error != nil {
		fmt.Fprintf(w, "There was an error")

		return
	}

	red, green, blue := sadboxColor.HexToRGB(sadboxColor.Hex(hexColor))

	colorSet := color.RGBA{
		R: red,
		G: green,
		B: blue,
	}

	dioderConfiguration.DioderInstance.SetAll(colorSet)
	fmt.Fprintf(w, message)
}
