package configurations

import "github.com/opentelekomcloud/gophertelekomcloud"

// List is used to obtain the parameter template list, including default
// parameter templates of all databases and those created by users.
func List(client *golangsdk.ServiceClient) (r ListResult) {
	raw, err := client.Get(client.ServiceURL("configurations"), nil, nil)
	return
}
