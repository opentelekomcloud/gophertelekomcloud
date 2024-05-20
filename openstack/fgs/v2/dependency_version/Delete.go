package dependency_version

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, dependId, version string) (err error) {
	_, err = client.Delete(client.ServiceURL("fgs", "dependencies", dependId, "version", version), nil)
	return
}
