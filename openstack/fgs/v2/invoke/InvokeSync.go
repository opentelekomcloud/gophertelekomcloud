package invoke

import (
	"net/http"

	"errors"
	"reflect"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// LaunchSync is used to execute a function synchronously.
// Clients must wait for explicit responses to their requests from the function.
// Responses are returned only after function invocation is complete.
func LaunchSync(client *golangsdk.ServiceClient, funcUrn string, body map[string]interface{}, headers LaunchSyncHeaders) (*LaunchSyncResp, *LaunchSyncResponseHeaders, error) {
	b, err := build.RequestBody(body, "")
	if err != nil {
		return nil, nil, err
	}

	// manually prepare headers and set defaults where needed
	headerMap := make(map[string]string)
	if headers.ContentType != "" {
		headerMap["Content-Type"] = headers.ContentType
	} else {
		// Content-Type header is mandatory
		return nil, nil, errors.New("Missing header 'Content-Type'.")
	}

	if headers.LogType != "" {
		headerMap["X-Cff-Log-Type"] = headers.LogType
	}

	// always set RequestVersion to "v1",
	// as response is expected in json format
	// v0 (text format) will not work here
	headerMap["X-Cff-Request-Version"] = "v1"

	if headers.InstanceMemory != "" {
		headerMap["X-Cff-Instance-Memory"] = headers.InstanceMemory
	}

	raw, err := client.Post(client.ServiceURL("fgs", "functions", funcUrn, "invocations"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: headerMap,
	})
	if err != nil {
		return nil, nil, err
	}

	resp_header := &LaunchSyncResponseHeaders{}
	unmarshalHeaders(raw.Header, resp_header)

	var res LaunchSyncResp
	err = extract.Into(raw.Body, &res)

	return &res, resp_header, err
}

// LaunchSyncHeaders represents the request headers for a synchronous function invocation
type LaunchSyncHeaders struct {
	// Content type of the request body.
	ContentType string `json:"Content-Type"` // e.g. "application/json"

	// Options: tail (4 KB logs will be returned) and null (no logs will be returned).
	LogType string `json:"X-Cff-Log-Type"`

	// Dynamic memory allocation for function.
	// Options:
	// - empty or
	// - any value of 128, 256, 512, 768, 1,024, 1,280, 1,536, 1,792,
	//   2,048, 2,560, 3,072, 3,584, 4,096, 8,192, or 10,240
	// If not empty, dynamic memory allocation must be enabled for the function.
	InstanceMemory string `json:"X-Cff-Instance-Memory"`
}

// NewLaunchSyncHeaders creates default headers for LaunchSync function invocation
func NewLaunchSyncHeaders() LaunchSyncHeaders {
	return LaunchSyncHeaders{
		ContentType:    "application/json",
		LogType:        "",
		InstanceMemory: "",
	}
}

// LaunchSyncResp represents the response from a synchronous function invocation
type LaunchSyncResp struct {
	// Request ID
	RequestID string `json:"request_id"`

	// Function execution result.
	Result string `json:"result"`

	// Function execution log.
	Log string `json:"log"`

	// Function execution status.
	Status int `json:"status"`
}

// LaunchSyncResponseHeaders represents the response headers from a synchronous function invocation
type LaunchSyncResponseHeaders struct {
	// Execution summary of the synchronous invocation.
	InvokeSummary string `json:"X-Cff-Invoke-Summary"`

	// Request ID of the synchronous invocation.
	RequestId string `json:"X-Cff-Request-Id"`

	// User log of the synchronous invocation.
	// Set X-Cff-Log-Type:tail in the request header.
	// Intercept and encode the last 2,000 bytes of the log using Base64.
	FunctionLog string `json:"X-Cff-Function-Log"`

	// Billing information of the synchronous invocation.
	BillingDuration string `json:"X-CFF-Billing-Duration"`

	// Response format:
	// - v0: text format
	// - v1: JSON format
	ResponseVersion string `json:"X-Cff-Response-Version"`

	// Error code of the synchronous invocation.
	// The value is 0 if the execution is successful.
	FuncErrCode string `json:"X-Func-Err-Code"`

	// Indicates whether the error occurs in a user function.
	IsFuncErr bool `json:"X-Is-Func-Err"`

	// AdditionalHeaders holds any additional headers not explicitly defined in the struct
	AdditionalHeaders map[string][]string `json:"-"`
}

// unmarshalHeaders unmarshals http.Header into a struct using json tags
func unmarshalHeaders(header http.Header, target interface{}) {
	val := reflect.ValueOf(target).Elem()
	typ := val.Type()

	// Track mapped headers
	mappedHeaders := make(map[string]bool)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		tag := fieldType.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}

		headerValue := header.Get(tag)
		if field.CanSet() && field.Kind() == reflect.String {
			field.SetString(headerValue)
			// Use http.CanonicalHeaderKey for case-insensitive comparison
			mappedHeaders[http.CanonicalHeaderKey(tag)] = true
		}
	}

	// Set AdditionalHeaders with unmapped headers only
	if field := val.FieldByName("AdditionalHeaders"); field.IsValid() && field.CanSet() {
		additionalHeaders := make(map[string][]string)
		for key, values := range header {
			canonicalKey := http.CanonicalHeaderKey(key)
			if !mappedHeaders[canonicalKey] {
				additionalHeaders[key] = values
			}
		}
		field.Set(reflect.ValueOf(additionalHeaders))
	}
}
