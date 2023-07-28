package golangsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

/*
Result is an internal type to be used by individual resource packages, but its
methods will be available on a wide variety of user-facing embedding types.

It acts as a base struct that other Result types, returned from request
functions, can embed for convenience. All Results capture basic information
from the HTTP transaction that was performed, including the response body,
HTTP headers, and any errors that happened.

Generally, each Result type will have an Extract method that can be used to
further interpret the result's payload in a specific context. Extensions or
providers can then provide additional extraction functions to pull out
provider- or extension-specific information as well.

Deprecated: use functions from internal/extract package instead
*/
type Result struct {
	// Body is the payload of the HTTP response from the server.
	Body []byte

	// Header contains the HTTP header structure from the original response.
	Header http.Header

	// Err is an error that occurred during the operation. It's deferred until
	// extraction to make it easier to chain the Extract call.
	Err error
}

type JsonRDSInstanceStatus struct {
	Instances  []JsonRDSInstanceField `json:"instances"`
	TotalCount int                    `json:"total_count"`
}

type JsonRDSInstanceField struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// Deprecated: use extract.Into function instead
func (r Result) ExtractInto(to any) error {
	if r.Err != nil {
		return r.Err
	}

	return extract.Into(bytes.NewReader(r.Body), to)
}

// Deprecated: use extract.IntoStructPtr function instead
func (r Result) ExtractIntoStructPtr(to any, label string) error {
	if r.Err != nil {
		return r.Err
	}

	return extract.IntoStructPtr(bytes.NewReader(r.Body), to, label)
}

// Deprecated: use extract.IntoSlicePtr function instead
func (r Result) ExtractIntoSlicePtr(to any, label string) error {
	if r.Err != nil {
		return r.Err
	}

	return extract.IntoSlicePtr(bytes.NewReader(r.Body), to, label)
}

// PrettyPrintJSON creates a string containing the full response body as
// pretty-printed JSON. It's useful for capturing test fixtures and for
// debugging extraction bugs. If you include its output in an issue related to
// a buggy extraction function, we will all love you forever.
func (r Result) String() string {
	pretty, err := json.MarshalIndent(r.Body, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	return string(pretty)
}

// ErrResult is an internal type to be used by individual resource packages, but
// its methods will be available on a wide variety of user-facing embedding
// types.
//
// It represents results that only contain a potential error and
// nothing else. Usually, if the operation executed successfully, the Err field
// will be nil; otherwise it will be stocked with a relevant error. Use the
// ExtractErr method
// to cleanly pull it out.
//
// Deprecated: use plain err return instead
type ErrResult struct {
	Err error
}

// ----------------------------------------------------------------------------

type ErrRespond struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

type ErrWithResult struct {
	ErrResult
}

func (r Result) Extract() (*ErrRespond, error) {
	var s = ErrRespond{}
	if r.Err != nil {
		return nil, r.Err
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, fmt.Errorf("failed to extract Error with Result")
	}
	return &s, nil
}

// ----------------------------------------------------------------------------

/*
HeaderResult is an internal type to be used by individual resource packages, but
its methods will be available on a wide variety of user-facing embedding types.

It represents a result that only contains an error (possibly nil) and an
http.Header. This is used, for example, by the objectstorage packages in
openstack, because most of the operations don't return response bodies, but do
have relevant information in headers.
*/
type HeaderResult struct {
	Result
}

// ExtractInto allows users to provide an object into which `Extract` will
// extract the http.Header headers of the result.
func (r HeaderResult) ExtractInto(to any) error {
	if r.Err != nil {
		return r.Err
	}

	tmpHeaderMap := map[string]string{}
	for k, v := range r.Header {
		if len(v) > 0 {
			tmpHeaderMap[k] = v[0]
		}
	}

	b, err := extract.JsonMarshal(tmpHeaderMap)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, to)

	return err
}

/*
Link is an internal type to be used in packages of collection resources that are
paginated in a certain way.

It's a response substructure common to many paginated collection results that is
used to point to related pages. Usually, the one we care about is the one with
Rel field set to "next".
*/
type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

/*
ExtractNextURL is an internal function useful for packages of collection
resources that are paginated in a certain way.

It attempts to extract the "next" URL from slice of Link structs, or
"" if no such URL is present.
*/
func ExtractNextURL(links []Link) (string, error) {
	var url string

	for _, l := range links {
		if l.Rel == "next" {
			url = l.Href
		}
	}

	if url == "" {
		return "", nil
	}

	return url, nil
}
