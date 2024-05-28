package json

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, code int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(value)

	if err != nil {
		return err
	}

	return nil
}
