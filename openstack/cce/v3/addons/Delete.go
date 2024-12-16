package addons

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete will permanently delete a particular addon based on its unique ID.
func Delete(client *golangsdk.ServiceClient, addonId, clusterId string) error {
	// DELETE /api/v3/addons/{id}?cluster_id={cluster_id}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("addons", addonId).WithQueryParams(&ClusterIdQueryParam{ClusterId: clusterId}).Build()
	if err != nil {
		return err
	}

	_, err = client.Delete(CCEServiceURL(client, clusterId, url.String()), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
		JSONBody:    nil,
	})
	if err != nil {
		return err
	}
	return nil
}
