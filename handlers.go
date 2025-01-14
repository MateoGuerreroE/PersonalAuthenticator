package main

import (
	"Personal2FA/totp"
	"Personal2FA/typings"
	"encoding/json"
	"net/http"
)

func HandleRegisterApp(w http.ResponseWriter, r *http.Request) {
	var requestBody typings.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if requestBody.AppName == "" || requestBody.AccountName == "" {
		http.Error(w, "Some of the required parameters is missing", http.StatusBadRequest)
	}

	var secret = totp.RegisterApp(requestBody.AppName, requestBody.AccountName)
	response := typings.ControllerResponse{Data: secret}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating JSON Response", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}

func HandleGenerateTOTP(w http.ResponseWriter, r *http.Request) {
	var localReq typings.GenerateRequest
	err := json.NewDecoder(r.Body).Decode(&localReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if localReq.AppName == "" {
		http.Error(w, "Some of the required parameters is missing", http.StatusBadRequest)
		return
	}

	code, totpErr := totp.GenerateTOTP(localReq.AppName)
	if totpErr != nil {
		http.Error(w, "Error Happened generating OTP", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, jsonErr := json.Marshal(typings.ControllerResponse{Data: code})
	if jsonErr != nil {
		http.Error(w, "Error generating JSON Response", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}
