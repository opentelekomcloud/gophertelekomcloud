package monitors

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular Health Monitor based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (*Monitor, error) {
	raw, err := client.Get(client.ServiceURL("healthmonitors", id), nil, nil)
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*Monitor, error) {
	if err != nil {
		return nil, err
	}

	var res Monitor
	err = extract.IntoStructPtr(raw.Body, res, "healthmonitor")
	return &res, err
}
