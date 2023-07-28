package golangsdk

import (
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
