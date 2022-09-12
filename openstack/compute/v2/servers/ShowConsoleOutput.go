package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// ShowConsoleOutput makes a request against the nova API to get console log from the server
func ShowConsoleOutput(client *golangsdk.ServiceClient, id string, opts ShowConsoleOutputOptsBuilder) (r ShowConsoleOutputResult) {
	b, err := opts.ToServerShowConsoleOutputMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("servers", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
