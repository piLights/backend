package webpage

import (
	"fmt"
	"net/http"
	"time"
)

func Logging(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		if dioderConfiguration.Debug {
			logChan <- fmt.Sprintf(
				"%s\t%s\t%s\t%s",
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
			)
		}
	})
}
