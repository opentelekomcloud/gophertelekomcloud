package organizations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetPermissions(client *golangsdk.ServiceClient, organization string) (*Permissions, error) {
	// PATCH /v2/manage/namespaces/{namespace}/access
	raw, err := client.Get(client.ServiceURL("manage", "namespaces", organization, "access"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Permissions
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Permissions struct {
	// Permission ID.
	ID int `json:"id"`
	// Organization name.
	Name string `json:"name"`
	// Organization creator.
	CreatorName string `json:"creator_name"`
	// Permissions of the current user.
	SelfAuth Auth `json:"self_auth"`
	// Permissions of other users.
	OthersAuth []Auth `json:"others_auths"`
}
