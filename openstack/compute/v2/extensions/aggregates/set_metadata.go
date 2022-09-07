package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// SetMetadata makes a request against the API to set metadata to a specific aggregate.
func SetMetadata(client *golangsdk.ServiceClient, aggregateID int, opts SetMetadataOpts) (*Aggregate, error) {
	b, err := opts.ToSetMetadataMap()
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
