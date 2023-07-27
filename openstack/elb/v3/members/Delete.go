package members

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete will remove and disassociate a Member from a particular
// Pool.
func Delete(client *golangsdk.ServiceClient, poolID string, memberID string) (err error) {
	// DELETE /v3/{project_id}/elb/pools/{pool_id}/members/{member_id}
	_, err = client.Delete(client.ServiceURL("pools", poolID, "members", memberID), nil)
	return
}
