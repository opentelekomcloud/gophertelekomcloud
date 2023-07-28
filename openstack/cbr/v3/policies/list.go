package policies

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"net/url"
)

// OperationType is a Policy type.
// One of `backup` and `replication`.
type OperationType string

type ListOpts struct {
	OperationType OperationType `q:"operation_type"`
	VaultID       string        `q:"vault_id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Policy, error) {
	var opts2 interface{} = opts
	query, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("policies")+query.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Policy
	err = extract.IntoSlicePtr(raw.Body, &res, "policies")
	return res, err
}
