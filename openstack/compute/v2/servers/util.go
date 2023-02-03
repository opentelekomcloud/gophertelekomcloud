package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// WaitForStatus will continually poll a server until it successfully
// transitions to a specified status. It will do this for at most the number of seconds specified.
func WaitForStatus(client *golangsdk.ServiceClient, id, status string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		current, err := Get(client, id)
		if err != nil {
			return false, err
		}

		if current.Status == status {
			return true, nil
		}

		return false, nil
	})
}

// IDFromName is a convenience function that returns a server's ID given its name.
func IDFromName(client *golangsdk.ServiceClient, name string) (string, error) {
	count := 0
	id := ""

	listOpts := ListOpts{
		Name: name,
	}

	allPages, err := List(client, listOpts).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractServers(allPages)
	if err != nil {
		return "", err
	}

	for _, f := range all {
		if f.Name == name {
			count++
			id = f.ID
		}
	}

	switch count {
	case 0:
		return "", golangsdk.ErrResourceNotFound{Name: name, ResourceType: "server"}
	case 1:
		return id, nil
	default:
		return "", golangsdk.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "server"}
	}
}
