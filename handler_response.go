package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type responseObj struct {
	Result  interface{} `json:"result"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

var Response = &response{}

type response struct {
}

func (r *response) SendGraphqlResponse(w http.ResponseWriter, msg, err string, status int, result interface{}) {
	if status != http.StatusOK {
		_ = Logger.LogError(msg, "handlers", "", fmt.Errorf(err))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(responseObj{Result: result, Error: err, Message: msg, Status: status})
}

// SendOkayResponse sends an Okay http response
func (r *response) SendOkayResponse(w http.ResponseWriter) error {
	return r.SendResponse(w, 200, map[string]string{})
}

// SendErrorResponse sends an Error http response
func (r *response) SendErrorResponse(w http.ResponseWriter, status int, message string) error {
	if status != http.StatusOK {
		_ = Logger.LogError(message, "handlers", "", nil)
	}
	return r.SendResponse(w, status, map[string]string{"error": message})
}

// SendResponse sends an http response
func (r *response) SendResponse(w http.ResponseWriter, status int, body interface{}) error {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(body)
}
