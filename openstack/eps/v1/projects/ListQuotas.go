package projects

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListQuotas(client *golangsdk.ServiceClient) (*Quotas, error) {
	raw, err := client.Get(client.ServiceURL("enterprise-projects", "quotas"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Quotas
	err = extract.IntoStructPtr(raw.Body, &res, "quotas")
	return &res, err
}

type Quotas struct {
	Resources []Quota `json:"resources"`
}

type Quota struct {
	Type  string `json:"type"`
	Used  int    `json:"used"`
	Quota int    `json:"quota"`
}
