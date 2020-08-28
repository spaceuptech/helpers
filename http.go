package helpers

import (
	"context"
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

func (r *response) SendGraphqlResponse(ctx context.Context, w http.ResponseWriter, msg, err string, status int, result interface{}) {
	if status != http.StatusOK {
		_ = Logger.LogError(GetRequestID(ctx), msg, fmt.Errorf(err), nil)
	}
	_ = r.SendResponse(ctx, w, http.StatusOK, responseObj{Result: result, Error: err, Message: msg, Status: status})
}

// SendOkayResponse sends an Okay http response
func (r *response) SendOkayResponse(ctx context.Context, w http.ResponseWriter) error {
	return r.SendResponse(ctx, w, 200, map[string]string{})
}

// SendErrorResponse sends an Error http response
func (r *response) SendErrorResponse(ctx context.Context, w http.ResponseWriter, status int, message string) error {
	if status != http.StatusOK {
		_ = Logger.LogError(GetRequestID(ctx), message, nil, nil)
	}
	return r.SendResponse(ctx, w, status, map[string]string{"error": message})
}

// SendResponse sends an http response
func (r *response) SendResponse(ctx context.Context, w http.ResponseWriter, status int, body interface{}) error {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	Logger.LogInfo(GetRequestID(ctx), "Response", map[string]interface{}{"statusCode": status})
	return json.NewEncoder(w).Encode(body)
}
