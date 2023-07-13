package pagination

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// PageResult stores the HTTP response that returned the current page of results.
type PageResult struct {
	golangsdk.Result
	url.URL
}

// GetBodyAsSlice tries to convert page body to a slice, returning nil on fail
func (r PageResult) GetBodyAsSlice() ([]interface{}, error) {
	result := make([]interface{}, 0)

	if err := extract.Into(bytes.NewReader(r.Body), &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetBodyAsMap tries to convert page body to a map, returning nil on fail
func (r PageResult) GetBodyAsMap() (map[string]interface{}, error) {
	result := make(map[string]interface{}, 0)

	if err := extract.Into(bytes.NewReader(r.Body), &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Request performs an HTTP request and extracts the http.Response from the result.
func Request(client *golangsdk.ServiceClient, headers map[string]string, url string) (*http.Response, error) {
	return client.Get(url, nil, &golangsdk.RequestOpts{
		MoreHeaders: headers,
		OkCodes:     []int{200, 204, 300},
	})
}
