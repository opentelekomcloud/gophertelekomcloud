package volumes

import "github.com/opentelekomcloud/gophertelekomcloud"

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

// IDFromName returns a volume ID for an exact name match.
func IDFromName(client *golangsdk.ServiceClient, name string) (string, error) {
	list, err := List(client, ListOpts{
		Name:  name,
		Limit: 1000,
	})
	if err != nil {
		return "", err
	}

	count := 0
	id := ""
	for _, volume := range list.Volumes {
		if volume.Name == name {
			count++
			id = volume.ID
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
