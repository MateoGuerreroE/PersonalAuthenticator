package handler

import (
	"encoding/json"
	"net/http"

	"Personal2FA/totp"
	"Personal2FA/typings"
)

func GenerateCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var localReq typings.GenerateRequest
	err := json.NewDecoder(r.Body).Decode(&localReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if localReq.AppName == "" {
		http.Error(w, "Some of the required parameters are missing", http.StatusBadRequest)
		return
	}

	code, totpErr := totp.GenerateTOTP(localReq.AppName)
	if totpErr != nil {
		http.Error(w, "Error happened generating OTP", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(typings.ControllerResponse{Data: code})
}
