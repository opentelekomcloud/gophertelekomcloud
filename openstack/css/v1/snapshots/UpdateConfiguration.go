package snapshots

import "github.com/opentelekomcloud/gophertelekomcloud"

func UpdateConfiguration(client *golangsdk.ServiceClient, clusterID string, opts UpdateConfigurationOptsBuilder) (r ErrorResult) {
	b, err := opts.ToUpdateConfigurationMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("clusters", clusterID, "index_snapshot", "setting"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
