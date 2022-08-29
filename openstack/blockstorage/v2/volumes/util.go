package volumes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// WaitForStatus will continually poll the resource, checking for a particular
// status. It will do this for the amount of seconds defined.
func WaitForStatus(c *golangsdk.ServiceClient, id, status string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		current, err := Get(c, id)
		if err != nil {
			return false, err
		}

		if current.Status == status {
			return true, nil
		}

		return false, nil
	})
}

// IDFromName is a convienience function that returns a server's ID given its name.
func IDFromName(client *golangsdk.ServiceClient, name string) (string, error) {
	count := 0
	id := ""

	listOpts := ListOpts{
		Name: name,
	}

	pages, err := List(client, listOpts).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractVolumes(pages)
	if err != nil {
		return "", err
	}

	for _, s := range all {
		if s.Name == name {
			count++
			id = s.ID
		}
	}

	switch count {
	case 0:
		return "", golangsdk.ErrResourceNotFound{Name: name, ResourceType: "volume"}
	case 1:
		return id, nil
	default:
		return "", golangsdk.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "volume"}
	}
}
