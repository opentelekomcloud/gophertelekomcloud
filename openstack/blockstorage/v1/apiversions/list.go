package apiversions

import (
	"net/url"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient) ([]APIVersion, error) {
	u, _ := url.Parse(client.ServiceURL(""))
	u.Path = "/"

	raw, err := client.Get(u.String(), nil, nil)
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
	// current status
	Status string `json:"status"`
	// date last updated
	Updated string `json:"updated"`
}
