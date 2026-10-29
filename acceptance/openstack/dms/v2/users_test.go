package v2

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v1/permissions"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/users"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUsersList(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := saslKafkaInstanceID(t, client)

	dmsUsers, err := users.List(client, instanceID)
	th.AssertNoErr(t, err)
	for _, val := range dmsUsers {
		tools.PrintResource(t, val)
	}
}

func TestUsersLifecycle(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)
	clientDmsV1, err := clients.NewDmsV11Client()
	th.AssertNoErr(t, err)

	instanceID := saslKafkaInstanceID(t, client)

	dmsTopic := createTopic(t, client, instanceID)
	t.Cleanup(func() { deleteTopic(t, client, instanceID, dmsTopic) })

	userName := tools.RandomString("user-", 5)
	userDesc := "acceptance test user"

	err = users.Create(client, instanceID, users.CreateOpts{
		UserName:   userName,
		UserDesc:   userDesc,
		UserPasswd: tools.RandomString("Aa1!", 12),
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err := users.Delete(client, instanceID, users.DeleteOpts{
			Action: "delete",
			Users:  []string{userName},
		})
		th.AssertNoErr(t, err)

		user, err := findUser(client, instanceID, userName)
		th.AssertNoErr(t, err)
		if user != nil {
			t.Errorf("user %s still exists after deletion", userName)
		}
	})

	user, err := findUser(client, instanceID, userName)
	th.AssertNoErr(t, err)
	if user == nil {
		t.Fatalf("user %s not found after creation", userName)
	}
	tools.PrintResource(t, user)
	th.AssertEquals(t, userDesc, user.UserDesc)
	if user.CreatedTime <= 0 {
		t.Errorf("unexpected created_time: %d", user.CreatedTime)
	}

	err = users.ResetPassword(client, instanceID, userName, users.ResetPasswordOpts{
		NewPassword: tools.RandomString("Bb2@", 12),
	})
	th.AssertNoErr(t, err)

	err = permissions.Create(clientDmsV1, instanceID, []permissions.CreateOpts{{
		Name: dmsTopic,
		Policies: []permissions.CreatePolicy{{
			UserName:     userName,
			AccessPolicy: "all",
		}},
	}})
	th.AssertNoErr(t, err)

	var policy *permissions.Policy
	err = golangsdk.WaitFor(60, func() (bool, error) {
		userPermissions, err := permissions.List(clientDmsV1, instanceID, dmsTopic)
		if err != nil {
			return false, err
		}
		th.AssertEquals(t, dmsTopic, userPermissions.Name)
		for i := range userPermissions.Policies {
			if userPermissions.Policies[i].UserName == userName {
				tools.PrintResource(t, userPermissions)
				policy = &userPermissions.Policies[i]
				return true, nil
			}
		}
		return false, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "all", policy.AccessPolicy)
}

func saslKafkaInstanceID(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Helper()

	instance := existingKafkaInstance(t, client, kafkaInstanceName(t))
	if !instance.SslEnable {
		t.Skipf("SASL is not enabled for Kafka instance %q, user management is unavailable", instance.Name)
	}
	return instance.InstanceID
}

func findUser(client *golangsdk.ServiceClient, instanceID, userName string) (*users.Users, error) {
	list, err := users.List(client, instanceID)
	if err != nil {
		return nil, err
	}
	for i := range list {
		if list[i].UserName == userName {
			return &list[i], nil
		}
	}
	return nil, nil
}
