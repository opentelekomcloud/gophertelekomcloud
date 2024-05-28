package async_config

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, funcURN string) (*AsyncInvokeResp, error) {
	raw, err := client.Get(client.ServiceURL("fgs", "functions", funcURN, "async-invoke-config"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res AsyncInvokeResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
