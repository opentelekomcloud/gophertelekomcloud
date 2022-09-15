package snapshots

import "github.com/opentelekomcloud/gophertelekomcloud"

func UpdateConfiguration(client *golangsdk.ServiceClient, clusterID string, opts UpdateConfigurationOptsBuilder) (r ErrorResult) {
	b, err := opts.ToUpdateConfigurationMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("clusters", clusterID, "index_snapshot", "setting"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
