package sqlfilter

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetSqlFilterSwitch(client *golangsdk.ServiceClient, instanceId string) (*SqlFilterSwitchResponse, error) {
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "sql-filter", "switch"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res SqlFilterSwitchResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type SqlFilterSwitchResponse struct {
	SwitchStatus string `json:"switch_status"`
}
