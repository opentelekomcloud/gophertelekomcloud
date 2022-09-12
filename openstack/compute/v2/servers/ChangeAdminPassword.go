package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// ChangeAdminPassword alters the administrator or root password for a specified server.
func ChangeAdminPassword(client *golangsdk.ServiceClient, id, newPassword string) (r ActionResult) {
	b := map[string]interface{}{
		"changePassword": map[string]string{
			"adminPass": newPassword,
		},
	}
	raw, err := client.Post(client.ServiceURL("servers", id, "action"), b, nil, nil)
	return
}
