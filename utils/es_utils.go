package utils

import (
	"encoding/json"
	"io"
)

func ParseESResponse[T any](body io.ReadCloser) (T, error) {
	defer body.Close()

	// Define a structure for parsing the ES response
	var esResponse struct {
		Found  bool `json:"found"`
		Source T    `json:"_source"`
	}

	if err := json.NewDecoder(body).Decode(&esResponse); err != nil {
		return *new(T), err
	}

	return esResponse.Source, nil
}
