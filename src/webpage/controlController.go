package webpage

import (
	"image/color"
	"net/http"

	"github.com/gorilla/mux"
)

func controlController(w http.ResponseWriter, r *http.Request) {
	state := mux.Vars(r)["state"]

	if state == "off" {
		dioderConfiguration.DioderInstance.SetAll(color.RGBA{})
	}
}
