package lifecycle

import "github.com/opentelekomcloud/gophertelekomcloud"

// UpdatePassword is updating password for a dcs instance
func UpdatePassword(client *golangsdk.ServiceClient, id string, opts UpdatePasswordOptsBuilder) (r UpdatePasswordResult) {

	body, err := opts.ToPasswordUpdateMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", id, "password"), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
