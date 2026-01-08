/*
Package http provides HTTP-related utility functions for the Flow library.

Author: Chisomo Chiweza (mwprogrammer)
*/
package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mwprogrammer/flow/internal/utilities/types"
)

// Response provides information about an HTTP request and response
type Response[T any] struct {
	URL          string
	Request      string
	Response     *string
	ResponseCode *int
	Data         *T
}

// Get sends an HTTP GET Request
func Get[T any](url string) (Response[T], error) {

	response := Response[T]{}

	response.URL = url
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

	defer func() {
		closeErr := content.Body.Close()
		if err == nil {
			err = closeErr
		}
	}()

	var result T

	if types.IsString(result) {
		return response, nil
	}

	err = json.Unmarshal(bytes, &result)

	if err != nil {
		return response, fmt.Errorf("error decoding JSON response: %w", err)
	}

	response.Data = &result

	return response, nil

}

// Post sends an HTTP POST request
func Post[T any, K any](url string, data T, headers map[string]string, timeout int) (Response[K], error) {

	response := Response[K]{}

	response.URL = url
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

	contentBody := string(bytes)

	response.Response = &contentBody

	defer func() {
		closeErr := content.Body.Close()
		if err == nil {
			err = closeErr
		}
	}()

	var result K

	if types.IsString(result) {
		return response, nil
	}

	err = json.Unmarshal(bytes, &result)

	if err != nil {
		return response, fmt.Errorf("error decoding JSON response: %w", err)
	}

	response.Data = &result

	return response, nil

}
