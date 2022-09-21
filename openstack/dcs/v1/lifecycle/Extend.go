package lifecycle

import "github.com/opentelekomcloud/gophertelekomcloud"

// Extend is extending for a dcs instance
func Extend(client *golangsdk.ServiceClient, id string, opts ExtendOptsBuilder) (r ExtendResult) {

	body, err := opts.ToExtendMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", id, "extend"), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
