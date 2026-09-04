package gpfs

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func cleanHeaderPrefix(header http.Header) map[string][]string {
	responseHeaders := make(map[string][]string)
	for key, value := range header {
		if len(value) > 0 {
			key = strings.ToLower(key)
			if strings.HasPrefix(key, HEADER_PREFIX) || strings.HasPrefix(key, HEADER_PREFIX_OBS) {
				key = key[len(HEADER_PREFIX):]
			}
			responseHeaders[key] = value
		}
	}
	return responseHeaders
}

// ConvertRequestToIoReader converts req to XML data
func ConvertRequestToIoReader(req interface{}) (io.Reader, error) {
	body, err := TransToXml(req)
	if err == nil {
		if isDebugLogEnabled() {
			doLog(LEVEL_DEBUG, "Do http request with data: %s", string(body))
		}
		return bytes.NewReader(body), nil
	}
	return nil, err
}

// ParseResponseToBaseModel gets response from OBS
func ParseResponseToBaseModel(resp *http.Response, baseModel IBaseModel, xmlResult bool, _ bool) (err error) {
	readCloser, ok := baseModel.(IReadCloser)
	if !ok {
		defer func() {
			errMsg := resp.Body.Close()
			if errMsg != nil {
				doLog(LEVEL_WARN, "Failed to close response body")
			}
		}()
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		if err == nil && len(body) > 0 {
			if xmlResult {
				err = ParseXml(body, baseModel)
			} else {
				err = parseJSON(body, baseModel)
			}
			if err != nil {
				doLog(LEVEL_ERROR, "Unmarshal error: %v", err)
			}
		}
	} else {
		readCloser.setReadCloser(resp.Body)
	}

	baseModel.setStatusCode(resp.StatusCode)
	responseHeaders := cleanHeaderPrefix(resp.Header)
	baseModel.setResponseHeaders(responseHeaders)
	if values, ok := responseHeaders[HEADER_REQUEST_ID]; ok {
		baseModel.setRequestId(values[0])
	}
	return
}

// ParseResponseToObsError gets obsError from OBS
func ParseResponseToObsError(resp *http.Response, isObs bool) error {
	obsError := ObsError{}
	respError := ParseResponseToBaseModel(resp, &obsError, true, isObs)
	if respError != nil {
		doLog(LEVEL_WARN, "Parse response to BaseModel with error: %v", respError)
	}
	obsError.Status = resp.Status
	return obsError
}
