package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mwprogrammer/flow/internal/utilities"
)

type HttpResponse[T any] struct {
	Url          string
	Request      string
	Response     *string
	ResponseCode *int
	Data         *T
}

func Get[T any](url string) (HttpResponse[T], error) {

	response := HttpResponse[T]{}

	response.Url = url
	response.Request = url

	content, err := http.Get(url)

	if err != nil {
		return response, err
	}

	response.ResponseCode = &content.StatusCode

	bytes, err := io.ReadAll(content.Body)

	if err != nil {
		return response, fmt.Errorf("failed to read response body: %w", err)
	}

	data := string(bytes)

	response.Response = &data

	defer content.Body.Close()

	if content.StatusCode != http.StatusCreated && content.StatusCode != http.StatusOK {

		error_bytes, _ := io.ReadAll(content.Body)
		return response, fmt.Errorf("received unexpected status code %d. Response: %s", content.StatusCode, error_bytes)

	}

	var result T

	if utilities.IsTypeString(result) {
		return response, nil
	}

	err = json.Unmarshal(bytes, &result)

	if err != nil {
		return response, fmt.Errorf("error decoding JSON response: %w", err)
	}

	response.Data = &result

	return response, nil

}

func Post[T any, K any](url string, data T, headers map[string]string, timeout int) (HttpResponse[K], error) {

	response := HttpResponse[K]{}

	response.Url = url
	response.Request = url

	body, err := json.Marshal(data)

	if err != nil {
		return response, fmt.Errorf("error marshaling request data: %w", err)
	}

	response.Request = string(body)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		return response, fmt.Errorf("error creating POST request: %w", err)
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	defer cancel()

	request = request.WithContext(ctx)

	client := &http.Client{}

	content, err := client.Do(request)

	if err != nil {
		return response, fmt.Errorf("error executing POST request: %w", err)
	}

	response.ResponseCode = &content.StatusCode

	bytes, err := io.ReadAll(content.Body)

	if err != nil {
		return response, fmt.Errorf("failed to read response body: %w", err)
	}

	content_body := string(bytes)

	response.Response = &content_body

	defer content.Body.Close()

	if content.StatusCode != http.StatusCreated && content.StatusCode != http.StatusOK {

		error_bytes, _ := io.ReadAll(content.Body)
		return response, fmt.Errorf("received unexpected status code %d. Response: %s", content.StatusCode, error_bytes)

	}

	var result K

	if utilities.IsTypeString(result) {
		return response, nil
	}

	err = json.Unmarshal(bytes, &result)

	if err != nil {
		return response, fmt.Errorf("error decoding JSON response: %w", err)
	}

	response.Data = &result

	return response, nil

}
