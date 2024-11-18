package clusters

import (
	"fmt"
	"io"
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Get(client *golangsdk.ServiceClient) (*string, error) {
	raw, err := client.Get(client.ServiceURL("cer", "download"), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "*/*",
		},
	})
	if err != nil {
		return nil, err
	}

	return ExtractCert(err, raw)
}

func ExtractCert(err error, response *http.Response) (*string, error) {
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	cert := string(bodyBytes)
	return &cert, nil
}
