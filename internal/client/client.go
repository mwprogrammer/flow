/*
Package client connects and interacts with the WhatsApp Cloud API.

Author: Chisomo Chiweza (mwprogrammer)
*/
package client

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mwprogrammer/flow/internal/utilities/http"
)

// ReadMessage reads and validates a JSON payload from Meta
func ReadMessage(payload string) (map[string]any, error) {

	var response BaseResponse

	err := json.Unmarshal([]byte(payload), &response)

	if err != nil {
		return nil, errors.New("the response body is not valid")
	}

	if response.Object == "" {
		return nil, errors.New("object property is not defined")
	}

	if response.Object != "whatsapp_business_account" {
		return nil, errors.New("object property is not whatsapp_business_account")
	}

	if len(response.Entries) == 0 {
		return nil, errors.New("entries property has empty array")
	}

	dataString, err := json.Marshal(response.Entries[0].Changes[0].Value)

	if err != nil {
		return nil, errors.New("value sub property could not be parsed")
	}

	var data map[string]any

	err = json.Unmarshal([]byte(dataString), &data)

	if err != nil {
		return nil, errors.New("value sub property is not valid")
	}

	return data, nil

}

// PostMessage sends a request to WhatsApp Cloud API
func PostMessage(version string, token string, sender string, payload any, endpoint string) error {

	headers := map[string]string{}

	headers["Content-Type"] = "application/json"
	headers["Accept-Language"] = "en_US"
	headers["Accept"] = "application"
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)

	baseURL := fmt.Sprintf("https://graph.facebook.com/v%s/%s", version, sender)
	url := fmt.Sprintf("%s/%s", baseURL, endpoint)

	_, err := http.Post[any, any](url, payload, headers, 30)

	if err != nil {
		return err
	}

	return nil

}
