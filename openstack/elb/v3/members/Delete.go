package members

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete will remove and disassociate a Member from a particular
// Pool.
func Delete(client *golangsdk.ServiceClient, poolID string, memberID string) (r DeleteResult) {
	_, r.Err = client.Delete(client.ServiceURL("pools", poolID, "members", memberID), nil)
	return
}
