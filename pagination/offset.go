package pagination

import (
	"fmt"
	"reflect"
	"strconv"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type OffsetPage interface {
	// LastElement returning index of the last element of the page
	LastElement() int
}

type OffsetPageBase struct {
	PageResult
}

func (p OffsetPageBase) LastElement() int {
	q := p.URL.Query()
	offset, err := strconv.Atoi(q.Get("offset"))
	if err != nil {
		panic(err)
	}
	limit, err := strconv.Atoi(q.Get("limit"))
	if err != nil {
		panic(err)
	}
	return offset + limit
}

func (p OffsetPageBase) NextPageURL() (string, error) {
	currentURL := p.URL
	q := currentURL.Query()
	if q.Get("offset") == "" && q.Get("limit") == "" {
		// without offset and limit it's just a SinglePageBase
		return "", nil
	}
	q.Set("offset", strconv.Itoa(p.LastElement()))
	currentURL.RawQuery = q.Encode()
	return currentURL.String(), nil
}

// IsEmpty returns true if this Page has no items in it.
func (p OffsetPageBase) IsEmpty() (bool, error) {
	if b, ok := p.Body.([]interface{}); ok {
		return len(b) == 0, nil
	}
	err := golangsdk.ErrUnexpectedType{}
	err.Expected = "[]interface{}"
	err.Actual = fmt.Sprintf("%v", reflect.TypeOf(p.Body))
	return true, err
}

// GetBody returns the Page Body. This is used in the `AllPages` method.
func (p OffsetPageBase) GetBody() interface{} {
	return p.Body
}
