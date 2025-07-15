package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type DeleteQueryParams struct {
	ErrorStatus string `q:"errorStatus,omitempty"`
	DeleteEfs   string `q:"delete_efs,omitempty"`
	DeleteENI   string `q:"delete_eni,omitempty"`
	DeleteEvs   string `q:"delete_evs,omitempty"`
	DeleteNet   string `q:"delete_net,omitempty"`
	DeleteObs   string `q:"delete_obs,omitempty"`
	DeleteSfs   string `q:"delete_sfs,omitempty"`
}

// Delete will permanently delete a particular cluster based on its unique ID.
func Delete(client *golangsdk.ServiceClient, clusterId string, opts DeleteQueryParams) error {

	url, err := golangsdk.NewURLBuilder().WithEndpoints("clusters", clusterId).WithQueryParams(opts).Build()
	if err != nil {
		return err
	}
	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
		JSONBody:    nil,
	})

	if err != nil {
		return err
	}
	return nil
}
