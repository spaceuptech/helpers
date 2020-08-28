package helpers

import (
	"context"
	"net/http"
)

type contextKey string

const contextKeyRequestID = contextKey("requestId")

func CreateContext(r *http.Request) context.Context {
	return context.WithValue(r.Context(), contextKeyRequestID, r.Header.Get(HeaderRequestID))
}

func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	value := ctx.Value(contextKeyRequestID)
	if value == nil {
		return noValueRequestID
	}
	return value.(string)
}
