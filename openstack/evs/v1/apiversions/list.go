package apiversions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient) ([]APIVersion, error) {
	raw, err := client.Get(client.ServiceURL(""), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []APIVersion
	err = extract.IntoSlicePtr(raw.Body, &res, "versions")
	return res, err
}

type APIVersion struct {
	// unique identifier
	ID string `json:"id"`
	// current status e.g. SUPPORTED
	Status string `json:"status"`
	// date last updated e.g. 2014-06-28T12:20:21Z
	Updated string `json:"updated"`
}
