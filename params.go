package golangsdk

import (
	"net/url"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

/*
Deprecated: use `internal/build.RequestBody` instead.
*/
func BuildRequestBody(opts interface{}, parent string) (map[string]interface{}, error) {
	body, err := build.RequestBody(opts, parent)
	if err != nil {
		return nil, err
	}

	return body.ToMap()
}

/*
Deprecated: use `internal/build.QueryString` instead.
*/
func BuildQueryString(opts interface{}) (*url.URL, error) {
	return build.QueryString(opts)
}

/*
Deprecated: use `internal/build.Headers` instead.
*/
func BuildHeaders(opts interface{}) (map[string]string, error) {
	return build.Headers(opts)
}
