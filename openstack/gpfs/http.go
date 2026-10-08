package gpfs

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (client Client) doActionWithoutFS(method string, input serializable, output responseModel) error {
	return client.doAction(method, "", input, output, true)
}

func (client Client) doActionWithFS(method, fsName string, input serializable, output responseModel) error {
	return client.doActionWithFSResult(method, fsName, input, output, true)
}

func (client Client) doActionWithFSResult(method, fsName string, input serializable, output responseModel, xmlResult bool) error {
	if strings.TrimSpace(fsName) == "" {
		return errors.New("file system name is empty")
	}
	return client.doAction(method, fsName, input, output, xmlResult)
}

func (client Client) doAction(method, fsName string, input serializable, output responseModel, xmlResult bool) error {
	params, headers, body, err := input.serialize()
	if err != nil {
		return err
	}
	if params == nil {
		params = make(map[string]string)
	}
	if headers == nil {
		headers = make(map[string][]string)
	}
	resp, err := client.doHTTP(method, strings.TrimSpace(fsName), params, headers, body)
	if err != nil {
		return err
	}
	if output == nil {
		_ = resp.Body.Close()
		return nil
	}
	return parseResponse(resp, output, xmlResult)
}

func requestBody(data interface{}) ([]byte, error) {
	switch value := data.(type) {
	case nil:
		return nil, nil
	case string:
		return []byte(value), nil
	case []byte:
		return value, nil
	default:
		return nil, fmt.Errorf("unsupported request body type %T", data)
	}
}

func (client Client) doHTTP(method, fsName string, params map[string]string, headers map[string][]string, data interface{}) (*http.Response, error) {
	body, err := requestBody(data)
	if err != nil {
		return nil, err
	}
	if body != nil {
		headers[headerContentLength] = []string{fmt.Sprintf("%d", len(body))}
	}

	var redirectURL string
	redirectCount := 0
	maxAttempts := client.conf.maxRetryCount + 1
	var lastErr error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		requestURL, err := client.authorize(method, fsName, params, headers, redirectHost(redirectURL))
		if err != nil {
			return nil, err
		}
		if redirectURL != "" {
			requestURL = redirectURL
		}

		req, err := http.NewRequest(method, requestURL, bytes.NewReader(body))
		if err != nil {
			return nil, fmt.Errorf("failed to build GPFS request: %w", err)
		}
		if client.conf.ctx != nil {
			req = req.WithContext(client.conf.ctx)
		}
		for key, values := range headers {
			switch {
			case strings.EqualFold(key, headerHost):
				if len(values) > 0 {
					req.Host = values[0]
				}
			case strings.EqualFold(key, headerContentLength):
				req.ContentLength = int64(len(body))
			default:
				req.Header[key] = append([]string(nil), values...)
			}
		}
		req.Header.Set(headerUserAgent, userAgent)
		resp, err := client.httpClient.Do(req)
		if err != nil {
			lastErr = err
		} else if resp.StatusCode < http.StatusMultipleChoices {
			return resp, nil
		} else if resp.StatusCode >= http.StatusBadRequest && resp.StatusCode < http.StatusInternalServerError {
			return nil, parseServiceError(resp)
		} else if resp.StatusCode >= http.StatusMultipleChoices && resp.StatusCode < http.StatusBadRequest {
			location := resp.Header.Get(headerLocation)
			if location == "" || redirectCount >= client.conf.maxRedirectCount {
				return nil, parseServiceError(resp)
			}
			_ = resp.Body.Close()
			redirectURL = location
			redirectCount++
			maxAttempts++
			continue
		} else {
			if attempt+1 >= maxAttempts {
				return nil, parseServiceError(resp)
			}
			lastErr = fmt.Errorf("GPFS service returned %s", resp.Status)
			_ = resp.Body.Close()
		}
		if attempt+1 < maxAttempts {
			time.Sleep(time.Duration(float64(attempt+1) * rand.Float64() * float64(time.Second)))
		}
	}
	return nil, fmt.Errorf("failed to send GPFS request: %w", lastErr)
}

func redirectHost(rawURL string) string {
	if rawURL == "" {
		return ""
	}
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return parsed.Host
}

type connDelegate struct {
	net.Conn
	socketTimeout time.Duration
	finalTimeout  time.Duration
}

func getConnDelegate(conn net.Conn, socketTimeout, finalTimeout int) *connDelegate {
	return &connDelegate{Conn: conn, socketTimeout: time.Duration(socketTimeout) * time.Second, finalTimeout: time.Duration(finalTimeout) * time.Second}
}

func (delegate *connDelegate) Read(data []byte) (int, error) {
	_ = delegate.SetReadDeadline(time.Now().Add(delegate.socketTimeout))
	n, err := delegate.Conn.Read(data)
	_ = delegate.SetReadDeadline(time.Now().Add(delegate.finalTimeout))
	return n, err
}

func (delegate *connDelegate) Write(data []byte) (int, error) {
	_ = delegate.SetWriteDeadline(time.Now().Add(delegate.socketTimeout))
	n, err := delegate.Conn.Write(data)
	deadline := time.Now().Add(delegate.finalTimeout)
	_ = delegate.SetWriteDeadline(deadline)
	_ = delegate.SetReadDeadline(deadline)
	return n, err
}
