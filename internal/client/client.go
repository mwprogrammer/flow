package client

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mwprogrammer/flow/internal/payloads"
	"github.com/mwprogrammer/flow/internal/utilities/http"
)

func ReadMessage(body string) (map[string]any, error) {

	var payload payloads.BaseResponse

	err := json.Unmarshal([]byte(body), &payload)

	if err != nil {
		return nil, errors.New("the response body is not valid")
	}

	if payload.Object == "" {
		return nil, errors.New("object property is not defined")
	}

	if payload.Object != "whatsapp_business_account" {
		return nil, errors.New("object property is not whatsapp_business_account")
	}

	if len(payload.Entries) == 0 {
		return nil, errors.New("entries property has empty array")
	}

	data_string, err := json.Marshal(payload.Entries[0].Changes[0].Value)

	if err != nil {
		return nil, errors.New("value sub property could not be parsed")
	}

	var data map[string]any

	err = json.Unmarshal([]byte(data_string), &data)

	if err != nil {
		return nil, errors.New("value sub property is not valid")
	}

	return data, nil

}

func PostMessage(version string, token string, sender string, payload any, endpoint string) error {

	headers := map[string]string{}

	headers["Content-Type"] = "application/json"
	headers["Accept-Language"] = "en_US"
	headers["Accept"] = "application"
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)

	base_url := fmt.Sprintf("https://graph.facebook.com/v%s/%s", version, sender)
	url := fmt.Sprintf("%s/%s", base_url, endpoint)

	_, err := http.Post[any, any](url, payload, headers, 30)

	if err != nil {
		return err
	}

	return nil

}
