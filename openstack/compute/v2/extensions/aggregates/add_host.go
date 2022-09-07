package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

type AddHostOpts struct {
	// The name of the host.
	Host string `json:"host" required:"true"`
}

// AddHost makes a request against the API to add host to a specific aggregate.
func AddHost(client *golangsdk.ServiceClient, aggregateID int, opts AddHostOpts) (r ActionResult) {
	b, err := golangsdk.BuildRequestBody(opts, "add_host")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("os-aggregates", strconv.Itoa(aggregateID), "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	return
}
