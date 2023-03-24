package main

import (
	"encoding/json"
	"net/http"
)

func main() {
    http.HandleFunc("/api/user", makeAPIFunc(handleUser))
    http.ListenAndServe(":3000", nil)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeAPIFunc(fn apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := fn(w, r); err != nil {
            writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        }
    }
}

func handleUser(w http.ResponseWriter, r *http.Request) error {
    return writeJSON(w, http.StatusOK, map[string]string{"message": "hello some user"})
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    return json.NewEncoder(w).Encode(v)
}
