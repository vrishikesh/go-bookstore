package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	var (
		body []byte
		err  error
	)

	if body, err = io.ReadAll(r.Body); err != nil {
		return fmt.Errorf("could not read: %w", err)
	}

	if err = json.Unmarshal(body, x); err != nil {
		return fmt.Errorf("could not unmarshal: %w", err)
	}

	return nil
}

func Json(w http.ResponseWriter, status int, data ...[]byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if len(data) > 0 {
		w.Write(data[0])
	}
}

func ErrorJson(w http.ResponseWriter, err error) {
	log.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
}
