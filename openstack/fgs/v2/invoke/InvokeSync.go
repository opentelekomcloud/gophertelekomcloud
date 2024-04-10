package invoke

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func LaunchSync(client *golangsdk.ServiceClient, funcUrn string) (*LaunchSyncResp, error) {
	raw, err := client.Post(client.ServiceURL("fgs", "functions", funcUrn, "invocations"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res LaunchSyncResp
	return &res, extract.Into(raw.Body, &res)
}

type LaunchSyncResp struct {
	RequestID string `json:"request_id"`
	Result    string `json:"result"`
	Log       string `json:"log"`
	Status    string `json:"status"`
}
