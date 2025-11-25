package client

import (
	"fmt"

	"github.com/mwprogrammer/flow/internal/http"
	"github.com/mwprogrammer/flow/models"
)

func PostMessage(settings models.Settings, payload models.Payload, endpoint string) error {

	headers := map[string]string{}

	headers["Content-Type"] = "application/json"
	headers["Accept-Language"] = "en_US"
	headers["Accept"] = "application"
	headers["Authorization"] = fmt.Sprintf("Bearer %s", settings.AccessToken)

	base_url := fmt.Sprintf("https://graph.facebook.com/%s/%s", settings.Version, settings.Phone)
	url := fmt.Sprintf("%s/%s", base_url, endpoint)

	_, err := http.Post[models.Payload, any](url, payload, headers, 30)

	if err != nil {
		return err
	}

	return nil

}
