package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register-app", HandleRegisterApp)
	mux.HandleFunc("POST /generate-code", HandleGenerateTOTP)

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", mux)
}
