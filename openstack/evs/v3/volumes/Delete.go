package volumes

import "github.com/opentelekomcloud/gophertelekomcloud"

// DeleteOpts contains options for deleting a Volume.
type DeleteOpts struct {
	VolumeId string
	// Specifies to delete all snapshots associated with the disk. The default value is false.
	Cascade bool `q:"cascade"`
}

// Delete will delete the existing Volume with the provided ID.
func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return
	}

	// DELETE /v3/{project_id}/volumes/{volume_id}
	_, err = client.Delete(client.ServiceURL("volumes", opts.VolumeId)+q.String(), nil)
	return
}
