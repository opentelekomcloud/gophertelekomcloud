package rocketmq

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/users"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const (
	userNotFound = "DMS.00500972"

	rocketMQUserSecret    = "5ecuredPa55w0rd!"
	rocketMQUserSecretNew = "2ecuredPa22w2rd!"
)

func TestRocketMQUsers(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := prepareRocketMQInstance(t, client)
	enableRocketMQACL(t, client, instanceID)

	topicName := createRocketMQTopic(t, client, instanceID)
	groupName := createRocketMQGroup(t, client, instanceID)

	// The service ignores admin, default_topic_perm, default_group_perm,
	// topic_perms and group_perms on both create and update: every user is
	// an administrator with PUB|SUB on topics and SUB on consumer groups.
	userName := tools.RandomString("user-acc-", 8)
	t.Logf("Attempting to create RocketMQ user: %s", userName)
	created, err := users.Create(client, instanceID, users.CreateOpts{
		AccessKey:          userName,
		SecretKey:          rocketMQUserSecret,
		WhiteRemoteAddress: "192.168.1.*",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, userName, created.AccessKey)

	t.Cleanup(func() {
		err := users.Delete(client, instanceID, userName)
		if err != nil && !hasErrorCode(err, userNotFound) {
			t.Errorf("failed to delete RocketMQ user %s: %s", userName, err)
		}
	})

	user, err := users.Get(client, instanceID, userName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, userName, user.AccessKey)
	th.AssertEquals(t, rocketMQUserSecret, user.SecretKey)
	th.AssertEquals(t, "192.168.1.*", user.WhiteRemoteAddress)
	th.AssertEquals(t, true, user.Admin)
	th.AssertEquals(t, "PUB|SUB", user.DefaultTopicPerm)
	th.AssertEquals(t, "SUB", user.DefaultGroupPerm)

	listed, err := users.List(client, instanceID, users.ListOpts{Limit: 50})
	th.AssertNoErr(t, err)
	var found bool
	for _, u := range listed.Users {
		if u.AccessKey == userName {
			found = true
			break
		}
	}
	th.AssertEquals(t, true, found)

	topicPolicies, err := users.ListTopicAccessPolicies(client, instanceID, topicName, users.ListAccessPoliciesOpts{Limit: 10})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, topicName, topicPolicies.Name)
	th.AssertEquals(t, "PUB|SUB", findAccessPolicy(t, topicPolicies.Policies, userName).Perm)

	groupPolicies, err := users.ListGroupAccessPolicies(client, instanceID, groupName, users.ListAccessPoliciesOpts{Limit: 10})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, groupName, groupPolicies.Name)
	th.AssertEquals(t, "SUB", findAccessPolicy(t, groupPolicies.Policies, userName).Perm)

	t.Logf("Attempting to update RocketMQ user: %s", userName)
	updated, err := users.Update(client, instanceID, userName, users.UpdateOpts{
		SecretKey:          rocketMQUserSecretNew,
		WhiteRemoteAddress: pointerto.String(""),
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, userName, updated.AccessKey)

	user, err = users.Get(client, instanceID, userName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, rocketMQUserSecretNew, user.SecretKey)
	th.AssertEquals(t, "", user.WhiteRemoteAddress)

	t.Logf("Attempting to delete RocketMQ user: %s", userName)
	th.AssertNoErr(t, users.Delete(client, instanceID, userName))
	_, err = users.Get(client, instanceID, userName)
	th.AssertEquals(t, true, hasErrorCode(err, userNotFound))
}

func findAccessPolicy(t *testing.T, policies []users.AccessPolicy, userName string) users.AccessPolicy {
	for _, p := range policies {
		if p.AccessKey == userName {
			return p
		}
	}
	t.Fatalf("no access policy for user %s", userName)
	return users.AccessPolicy{}
}

// enableRocketMQACL turns on access control, which user management requires,
// and restores the previous setting on cleanup.
func enableRocketMQACL(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	instance, err := instances.Get(client, instanceID)
	th.AssertNoErr(t, err)
	if instance.EnableACL {
		return
	}

	t.Logf("Attempting to enable ACL of RocketMQ instance: %s", instanceID)
	th.AssertNoErr(t, updateRocketMQACL(client, instanceID, true))

	t.Cleanup(func() {
		t.Logf("Attempting to disable ACL of RocketMQ instance: %s", instanceID)
		th.AssertNoErr(t, updateRocketMQACL(client, instanceID, false))
	})
}

func updateRocketMQACL(client *golangsdk.ServiceClient, instanceID string, enable bool) error {
	err := instances.Update(client, instanceID, instances.UpdateOpts{EnableACL: pointerto.Bool(enable)})
	if err != nil {
		return err
	}

	return golangsdk.WaitFor(600, func() (bool, error) {
		instance, err := instances.Get(client, instanceID)
		if err != nil {
			return false, err
		}
		return instance.EnableACL == enable && instance.Status == rocketMQTargetStatus, nil
	})
}
