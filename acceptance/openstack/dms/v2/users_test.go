package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v1/permissions"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/users"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUsersList(t *testing.T) {
	t.Skip("DMS Creation takes too long to complete")
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := createDmsInstance(t, client)
	defer deleteDmsInstance(t, client, instanceID)

	dmsUsers, err := users.List(client, instanceID)
	th.AssertNoErr(t, err)
	for _, val := range dmsUsers {
		tools.PrintResource(t, val)
	}
}

func TestUsersLifecycle(t *testing.T) {
	t.Skip("DMS Creation takes too long to complete")
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)
	clientDmsV1, err := clients.NewDmsV11Client()
	th.AssertNoErr(t, err)

	instanceID := createDmsInstance(t, client)
	defer deleteDmsInstance(t, client, instanceID)

	dmsTopic := createTopic(t, client, instanceID)
	defer deleteTopic(t, client, instanceID, dmsTopic)

	userName := tools.RandomString("user", 5)

	createOpts := users.CreateOpts{
		UserName:   userName,
		UserPasswd: "test12312!",
	}
	err = users.Create(client, instanceID, createOpts)
	th.AssertNoErr(t, err)

	err = users.ResetPassword(client, instanceID, userName, "tes!$%t12312")
	th.AssertNoErr(t, err)

	err = permissions.Create(clientDmsV1, instanceID, []permissions.CreateOpts{{
		Name: dmsTopic,
		Policies: []permissions.CreatePolicy{
			{
				UserName:     userName,
				AccessPolicy: "all",
			},
		},
	},
	})

	th.AssertNoErr(t, err)

	userPermissions, err := permissions.List(clientDmsV1, instanceID, dmsTopic)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, userPermissions)

	err = users.Delete(client, instanceID, users.DeleteOpts{
		Action: "delete",
		Users: []string{
			userName,
		},
	})
	th.AssertNoErr(t, err)
}
