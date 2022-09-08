package volumes

import "github.com/opentelekomcloud/gophertelekomcloud"

type DeleteOpts struct {
	// Specifies to delete all snapshots associated with the EVS disk.
	Cascade bool `q:"cascade"`
}

func Delete(client *golangsdk.ServiceClient, id string, opts DeleteOpts) (err error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL("volumes", id)+q.String(), nil)
	return
}
