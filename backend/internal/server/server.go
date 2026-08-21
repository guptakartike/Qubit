package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Qubit is running")
}

func healthHandler(w http.ResponseWriter, r *http.Request){
	response:= map[string]string{
		"status":"ok",
		"service":"qubit-api",
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func New(port int) *http.Server{
	mux:=http.NewServeMux()
	mux.HandleFunc("/api/v1/",homeHandler)
	mux.HandleFunc("/api/v1/health",healthHandler)

	return &http.Server{
		Addr:fmt.Sprintf(":%d",port),
		Handler:mux,
	}

	

	
}