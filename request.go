package xhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DecodeRequest(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("failed to read body from %v", r)
	}

	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal request as type %t: %v", v, err)
	}
	return nil
}
