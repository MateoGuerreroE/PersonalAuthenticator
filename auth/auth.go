package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// IsAuthorized checks the Authorization header against APP_PASS.
// Returns (allowed, error). If APP_PASS is not configured, returns an error
// so callers can respond with 500 (server misconfiguration) instead of 401.
// Accepts either a raw token or "Bearer <token>".
func IsAuthorized(r *http.Request) (bool, error) {
	appPass := strings.TrimSpace(os.Getenv("APP_PASS"))
	if appPass == "" {
		return false, fmt.Errorf("APP_PASS not set")
	}

	authHeader := strings.TrimSpace(r.Header.Get("Authorization"))
	if authHeader == "" {
		return false, nil
	}

	if authHeader == appPass {
		return true, nil
	}

	bearerValue := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	return bearerValue == appPass, nil
}
