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

// EnabledState is a convenience type, mostly used in Create and Update
// operations. Because the zero value of a bool is FALSE, we need to use a
// pointer instead to indicate zero-ness.
// Deprecated, use pointerto.Bool instead
type EnabledState *bool

// Convenience vars for EnabledState values.
// Deprecated: use `pointerto.Bool` instead.
var (
	iTrue  = true
	iFalse = false

	// Enabled is a pointer to `true`.
	Enabled EnabledState = &iTrue
	// Disabled is a pointer to `false`.
	Disabled EnabledState = &iFalse
)

// IPVersion is a type for the possible IP address versions. Valid instances
// are IPv4 and IPv6
type IPVersion int

const (
	// IPv4 is used for IP version 4 addresses
	IPv4 IPVersion = 4
	// IPv6 is used for IP version 6 addresses
	IPv6 IPVersion = 6
)

var t time.Time

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
		if v.Type() == reflect.TypeOf(t) {
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
	optsValue := reflect.ValueOf(opts)
	if optsValue.Kind() == reflect.Ptr {
		optsValue = optsValue.Elem()
	}

	optsType := reflect.TypeOf(opts)
	if optsType.Kind() == reflect.Ptr {
		optsType = optsType.Elem()
	}

	params := url.Values{}

	if optsValue.Kind() == reflect.Struct {
		for i := 0; i < optsValue.NumField(); i++ {
			v := optsValue.Field(i)
			f := optsType.Field(i)
			qTag := f.Tag.Get("q")

			// if the field has a 'q' tag, it goes in the query string
			if qTag != "" {
				tags := strings.Split(qTag, ",")

				// if the field is set, add it to the slice of query pieces
				if !isZero(v) {
				loop:
					switch v.Kind() {
					case reflect.Ptr:
						v = v.Elem()
						goto loop
					case reflect.String:
						params.Add(tags[0], v.String())
					case reflect.Int:
						params.Add(tags[0], strconv.FormatInt(v.Int(), 10))
					case reflect.Int64:
						params.Add(tags[0], strconv.FormatInt(v.Int(), 10))
					case reflect.Bool:
						params.Add(tags[0], strconv.FormatBool(v.Bool()))
					case reflect.Slice:
						switch v.Type().Elem() {
						case reflect.TypeOf(0):
							for i := 0; i < v.Len(); i++ {
								params.Add(tags[0], strconv.FormatInt(v.Index(i).Int(), 10))
							}
						default:
							for i := 0; i < v.Len(); i++ {
								params.Add(tags[0], v.Index(i).String())
							}
						}
					case reflect.Map:
						if v.Type().Key().Kind() == reflect.String && v.Type().Elem().Kind() == reflect.String {
							var s []string
							for _, k := range v.MapKeys() {
								value := v.MapIndex(k).String()
								s = append(s, fmt.Sprintf("'%s':'%s'", k.String(), value))
							}
							params.Add(tags[0], fmt.Sprintf("{%s}", strings.Join(s, ", ")))
						}
					}
				} else {
					// Otherwise, the field is not set.
					if len(tags) == 2 && tags[1] == "required" {
						// And the field is required. Return an error.
						return &url.URL{}, fmt.Errorf("required query parameter [%s] not set", f.Name)
					}
				}
			}
		}

		return &url.URL{RawQuery: params.Encode()}, nil
	}
	// Return an error if the underlying type of 'opts' isn't a struct.
	return nil, fmt.Errorf("options type is not a struct")
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
