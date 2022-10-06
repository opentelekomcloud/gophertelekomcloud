package policies

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// OperationType is a Policy type.
// One of `backup` and `replication`.
type OperationType string

type ListOpts struct {
	OperationType OperationType `q:"operation_type"`
	VaultID       string        `q:"vault_id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Policy, error) {
	query, err := golangsdk.BuildQueryString(opts)
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
