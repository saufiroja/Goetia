package middlewares

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()

		// Call the next handler in the chain
		next(w, r, ps)

		// Log information about the request
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	}
}
