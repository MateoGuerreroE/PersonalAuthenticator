package handler

import (
	"net/http"
	"os"
	"strings"
)

// authorizeRequest validates Authorization against APP_PASS.
// It accepts either a raw token value or "Bearer <token>".
func authorizeRequest(r *http.Request) bool {
	appPass := strings.TrimSpace(os.Getenv("APP_PASS"))
	if appPass == "" {
		return false
	}

	authHeader := strings.TrimSpace(r.Header.Get("Authorization"))
	if authHeader == "" {
		return false
	}

	if authHeader == appPass {
		return true
	}

	bearerValue := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	return bearerValue == appPass
}
