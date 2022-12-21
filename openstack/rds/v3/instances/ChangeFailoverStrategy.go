package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ChangeFailoverStrategyOpts struct {
	InstanceId string `json:"-"`
	// Specifies the failover priority. Valid value:
	// reliability: Data reliability is preferentially ensured during the failover to minimize the amount of lost data. It is recommended for services that require high data consistency.
	// availability: Data availability is preferentially ensured during the failover to recover services quickly. It is recommended for services that have high requirements on the database online duration.
	RepairStrategy string `json:"repairStrategy"`
}

func ChangeFailoverStrategy(client *golangsdk.ServiceClient, opts ChangeFailoverStrategyOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT /v3/{project_id}/instances/{instance_id}/failover/strategy
	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId, "failover", "strategy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
