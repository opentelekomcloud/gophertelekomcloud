package gpfs

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func responseHeaders(header http.Header) map[string][]string {
	result := make(map[string][]string, len(header))
	for key, values := range header {
		name := strings.ToLower(key)
		name = strings.TrimPrefix(name, headerPrefixOBS)
		result[name] = append([]string(nil), values...)
	}
	return result
}

func parseResponse(resp *http.Response, model responseModel, xmlResult bool) error {
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read GPFS response: %w", err)
	}
	if len(body) > 0 {
		if xmlResult {
			err = parseXML(body, model)
		} else {
			err = parseJSON(body, model)
		}
		if err != nil {
			return fmt.Errorf("failed to decode GPFS response: %w", err)
		}
	}
	model.setStatusCode(resp.StatusCode)
	headers := responseHeaders(resp.Header)
	model.setResponseHeaders(headers)
	if values := headers[headerRequestID]; len(values) > 0 {
		model.setRequestID(values[0])
	}
	return nil
}

func parseServiceError(resp *http.Response) error {
	serviceError := Error{}
	if err := parseResponse(resp, &serviceError, true); err != nil {
		return fmt.Errorf("gpfs: failed to parse error response with status %s: %w", resp.Status, err)
	}
	serviceError.Status = resp.Status
	return serviceError
}
