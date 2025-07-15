package dependency_version

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, dependId, version string) (*DepVersionResp, error) {
	raw, err := client.Get(client.ServiceURL("fgs", "dependencies", dependId, "version", version), nil, nil)
	if err != nil {
		return nil, err
	}

	var res DepVersionResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
