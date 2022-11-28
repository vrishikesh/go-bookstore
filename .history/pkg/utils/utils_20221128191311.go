package utils

import (
	"encoding/json"
	"fmt"
	"io"
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