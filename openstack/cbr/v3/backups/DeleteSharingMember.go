package backups

import "github.com/opentelekomcloud/gophertelekomcloud"

func DeleteSharingMember(client *golangsdk.ServiceClient, id, memberID string) (err error) {
	_, err = client.Delete(client.ServiceURL("backups", id, "members", memberID), nil)
	return
}
