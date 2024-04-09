package function

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetFuncMetadata is basically /GET instance function
func GetMetadata(client *golangsdk.ServiceClient, funcURN string) (*FuncGraph, error) {
	raw, err := client.Get(client.ServiceURL("fgs", "functions", funcURN, "config"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res FuncGraph
	err = extract.Into(raw.Body, &res)
	return &res, err
}
