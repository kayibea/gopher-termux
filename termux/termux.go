package termux

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
)

type ApiError struct {
	Error    string `json:"error"`
	APIError string `json:"API_ERROR"`
}

// Keep thinking, maybe this should be an internal

func Exec(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	var apiErr ApiError
	output := out.Bytes()

	if err := json.Unmarshal(output, &apiErr); err != nil {
		return nil, err
	}

	if apiErr.Error != "" {
		return nil, errors.New(apiErr.Error)
	}

	if apiErr.APIError != "" {
		return nil, errors.New(apiErr.APIError)
	}

	// var apiErr ApiError
	// if err := json.Unmarshal(output, &apiErr); err == nil && apiErr.APIError != "" {
	// 	// return nil, errors.New(apiErr.APIError)
	// }

	return output, nil
}
