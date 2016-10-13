package webpage

import (
	"net/http"

	"gitlab.com/piLights/dioder-rpc/src/configuration"
)

var (
	logChan             chan interface{}
	fatalChan           chan interface{}
	dioderConfiguration configuration.Configuration
)

func StartWebPage(logCh, fatalCh chan interface{}, config configuration.Configuration) {
	logChan = logCh
	fatalChan = fatalCh

	dioderConfiguration = config

	if dioderConfiguration.Debug {
		logChan <- "Starting REST-API"
		logChan <- "Registering available routes"
	}

	colorAPI := NewRouter()

	fatalChan <- http.ListenAndServe(":8080", colorAPI)
}
