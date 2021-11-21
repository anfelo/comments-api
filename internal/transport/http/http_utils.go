package http

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Error(err)
	}
}
