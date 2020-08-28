package helpers

import (
	"context"
	"encoding/json"
	"net/http"
)

var Response = &response{}

type response struct {
}

// SendOkayResponse sends an Okay http response
func (r *response) SendOkayResponse(ctx context.Context, statusCode int, w http.ResponseWriter) error {
	return r.SendResponse(ctx, w, statusCode, map[string]string{})
}

// SendErrorResponse sends an Error http response
func (r *response) SendErrorResponse(ctx context.Context, w http.ResponseWriter, statusCode int, message string) error {
	if statusCode != http.StatusOK {
		_ = Logger.LogError(GetRequestID(ctx), message, nil, nil)
	}
	return r.SendResponse(ctx, w, statusCode, map[string]string{"error": message})
}

// SendResponse sends an http response
func (r *response) SendResponse(ctx context.Context, w http.ResponseWriter, statusCode int, body interface{}) error {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	Logger.LogInfo(GetRequestID(ctx), "Response", map[string]interface{}{"statusCode": statusCode})
	return json.NewEncoder(w).Encode(body)
}
