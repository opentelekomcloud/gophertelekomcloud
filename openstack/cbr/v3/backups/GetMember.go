package backups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetSharingMember(client *golangsdk.ServiceClient, id, memberID string) (*Member, error) {
	raw, err := client.Get(client.ServiceURL("backups", id, "members", memberID), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Member
	err = extract.IntoStructPtr(raw.Body, &res, "member")
	return &res, err
}
