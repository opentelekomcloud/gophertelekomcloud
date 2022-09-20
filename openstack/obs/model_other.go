package obs

import (
	"net/http"
)

// CreateSignedUrlInput is the input parameter of CreateSignedUrl function
type CreateSignedUrlInput struct {
	Method      HttpMethodType
	Bucket      string
	Key         string
	SubResource SubResourceType
	Expires     int
	Headers     map[string]string
	QueryParams map[string]string
}

// CreateSignedUrlOutput is the result of CreateSignedUrl function
type CreateSignedUrlOutput struct {
	SignedUrl                  string
	ActualSignedRequestHeaders http.Header
}

// CreateBrowserBasedSignatureInput is the input parameter of CreateBrowserBasedSignature function.
type CreateBrowserBasedSignatureInput struct {
	Bucket     string
	Key        string
	Expires    int
	FormParams map[string]string
}

// CreateBrowserBasedSignatureOutput is the result of CreateBrowserBasedSignature function.
type CreateBrowserBasedSignatureOutput struct {
	OriginPolicy string
	Policy       string
	Algorithm    string
	Credential   string
	Date         string
	Signature    string
}
