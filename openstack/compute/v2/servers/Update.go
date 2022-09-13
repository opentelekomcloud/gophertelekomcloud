package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// Update requests that various attributes of the indicated server be changed.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (*Server, error) {
	b, err := opts.ToServerUpdateMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("servers", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return ExtractSer(err, raw)
}
