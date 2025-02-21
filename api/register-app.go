package handler

import (
	"encoding/json"
	"net/http"

	"Personal2FA/totp"
	"Personal2FA/typings"
)

func RegisterAppHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody typings.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if requestBody.AppName == "" || requestBody.AccountName == "" {
		http.Error(w, "Some of the required parameters are missing", http.StatusBadRequest)
		return
	}

	totp.RegisterApp(requestBody.AppName, requestBody.AccountName, requestBody.Secret)
	response := typings.ControllerResponse{Data: "App Registered"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
