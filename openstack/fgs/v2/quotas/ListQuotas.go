package quotas

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListQuotas(client *golangsdk.ServiceClient) (*ListQuotasResults, error) {
	raw, err := client.Get(client.ServiceURL("fgs", "quotas"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListQuotasResults
	err = extract.IntoStructPtr(raw.Body, &res, "quotas")
	return &res, err
}

type ListQuotasResults struct {
	Resources []Resources `json:"resources"`
}

type Resources struct {
	Quota int    `json:"quota"`
	Used  int    `json:"used"`
	Type  string `json:"type"`
	Unit  string `json:"unit"`
}
