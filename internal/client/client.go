package client

import (
	"fmt"

	"github.com/mwprogrammer/flow/internal/utilities/http"
)

func ReadMessage(request string) {}

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
