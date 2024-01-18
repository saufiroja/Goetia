package requests

import (
	"encoding/json"
	"net/http"
)

func NewRequestMapper(r *http.Request, input interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return err
	}

	return nil
}
