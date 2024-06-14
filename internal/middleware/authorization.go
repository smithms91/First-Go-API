package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/smithms91/goapi/api"
	"github.com/smithms91/goapi/internal/tools"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}

		loginDetails := (*database).GetLoginDetails(username)
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
