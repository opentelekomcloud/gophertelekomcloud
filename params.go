package golangsdk

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

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

	var res map[string]interface{}
	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(marshal, &res)
	return res, err
}

// isZero checks if given argument has default type value.
func isZero(v reflect.Value) bool {
	// fmt.Printf("\n\nchecking isZero for value: %+v\n", v)
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return false
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		var warp time.Time
		if v.Type() == reflect.TypeOf(warp) {
			return v.Interface().(time.Time).IsZero()
		}
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && isZero(v.Field(i))
		}
		return z
	}
	// Compare other types directly:
	z := reflect.Zero(v.Type())
	// fmt.Printf("zero type for value: %+v\n\n\n", z)
	return v.Interface() == z.Interface()
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
	optsValue := reflect.ValueOf(opts)
	if optsValue.Kind() == reflect.Ptr {
		optsValue = optsValue.Elem()
	}

	optsType := reflect.TypeOf(opts)
	if optsType.Kind() == reflect.Ptr {
		optsType = optsType.Elem()
	}

	optsMap := make(map[string]string)
	if optsValue.Kind() == reflect.Struct {
		for i := 0; i < optsValue.NumField(); i++ {
			v := optsValue.Field(i)
			f := optsType.Field(i)
			hTag := f.Tag.Get("h")

			// if the field has a 'h' tag, it goes in the header
			if hTag != "" {
				tags := strings.Split(hTag, ",")

				// if the field is set, add it to the slice of query pieces
				if !isZero(v) {
					switch v.Kind() {
					case reflect.String:
						optsMap[tags[0]] = v.String()
					case reflect.Int:
						optsMap[tags[0]] = strconv.FormatInt(v.Int(), 10)
					case reflect.Bool:
						optsMap[tags[0]] = strconv.FormatBool(v.Bool())
					}
				} else {
					// Otherwise, the field is not set.
					if len(tags) == 2 && tags[1] == "required" {
						// And the field is required. Return an error.
						return optsMap, fmt.Errorf("Required header not set.")
					}
				}
			}

		}
		return optsMap, nil
	}
	// Return an error if the underlying type of 'opts' isn't a struct.
	return optsMap, fmt.Errorf("options type is not a struct")
}
