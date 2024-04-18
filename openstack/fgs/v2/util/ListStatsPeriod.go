package util

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListStatsPeriod(client *golangsdk.ServiceClient, funcURN, period string) (*StatResp, error) {
	raw, err := client.Get(client.ServiceURL("fgs", "functions", funcURN, "statistics", period), nil, nil)
	if err != nil {
		return nil, err
	}

	var res StatResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
