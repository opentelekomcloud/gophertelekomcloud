package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type AddHostOpts struct {
	// The name of the host.
	Host string `json:"host" required:"true"`
}

// AddHost makes a request against the API to add host to a specific aggregate.
func AddHost(client *golangsdk.ServiceClient, aggregateID int, opts AddHostOpts) (*Aggregate, error) {
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

	var res Aggregate
	err = extract.IntoStructPtr(raw.Body, &res, "aggregate")
	return &res, err
}
